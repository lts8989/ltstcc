package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"log"
	"ltstcc/model"
	"ltstcc/users"
	"ltstcc/utility"
	"net"
)

type server struct {
	usersMap map[string]*users.Users
	users.UnimplementedUsersInfoServer
}

func (s *server) CreateUsers(ctx context.Context, req *users.Users) (resp *users.UsersId, err error) {
	resp = &users.UsersId{}
	var users = model.Users{
		Username:     req.Username,
		Balance:      req.Balance,
		FrozenAmount: req.FrozenAmount,
	}
	err = utility.GetUsersDb().Create(&users).Error
	if err != nil {
		fmt.Println("创建失败")
	}
	fmt.Printf("创建成功,userId:%d", users.ID)
	resp.Value = int32(users.ID)
	return
}

func (s *server) BatchCreateUsers(ctx context.Context, req *users.BatchCreateUsersParams) (resp *users.BatchCreateUsersCount, err error) {
	resp = &users.BatchCreateUsersCount{}
	var buffer bytes.Buffer
	sql := "insert into `users` (`username`,`balance`,`frozen_amount`) values"
	if _, err = buffer.WriteString(sql); err != nil {
		return
	}

	for i := 0; i < int(req.Count); i++ {
		buffer.WriteString(fmt.Sprintf("('%s','%d','%d')", utility.RandString(10), 0, 0))
		if i == (int(req.Count) - 1) {
			buffer.WriteString(";")
		} else {
			buffer.WriteString(",")
		}
	}
	db := utility.GetUsersDb().Exec(buffer.String())
	err = db.Error
	resp.Count = db.RowsAffected
	return
}

func (s *server) Recharge(ctx context.Context, req *users.RechargeParams) (resp *users.UsersSuccess, err error) {
	resp = &users.UsersSuccess{}
	user, err := getUsersInfo(req.UserId)
	if err != nil {
		return
	}

	err = utility.GetUsersDb().Transaction(func(tx *gorm.DB) error {

		err = tx.Model(model.Users{}).Where("id=?", req.UserId).Update("balance", gorm.Expr("balance+?", req.RechargeMoney)).Error
		if err != nil {
			fmt.Println("user balance更新失败")
			return err
		}

		var usersDetail = model.UsersDetail{
			UserId:       req.UserId,
			Balance:      user.Balance + req.RechargeMoney,
			FrozenAmount: user.FrozenAmount,
			Type:         utility.USERS_DETAIL_TYPE_RECHARGE,
		}
		err = tx.Create(&usersDetail).Error
		if err != nil {
			fmt.Println("users_detail 插入失败")
			return err
		}

		return nil
	})

	if err != nil {
		resp.Value = 0
		return
	}
	resp.Value = 1

	return
}

func getUsersInfo(userId int32) (model.Users, error) {
	var user = model.Users{}
	err := utility.GetUsersDb().First(&user, userId).Error
	if err != nil {
		fmt.Println("查找用户失败", err)
	}
	return user, err
}

func lockUser(usersId int32) (err error) {
	key := fmt.Sprintf(utility.USERS_LOCK_KEY, usersId)
	val, err := utility.GetRedis().Get(key).Result()
	if val != "" {
		err = errors.New("用户已经加锁")
		return
	}
	err = utility.GetRedis().Set(key, utility.USERS_DETAIL_TYPE_RECHARGE, 0).Err()
	return
}

func (s *server) FrozenAmount(ctx context.Context, req *users.FrozenTryParams) (resp *users.FrozenTryResp, err error) {
	resp = &users.FrozenTryResp{}

	user, err := getUsersInfo(req.UserId)
	if err != nil {
		return
	}

	if user.Balance < req.FrozenAmount {
		err = errors.New(fmt.Sprintf("用户余额不足，冻结金额失败，当前余额：%d,冻结金额:%d", user.Balance, req.FrozenAmount))
		return
	}

	resp.UserBalance = user.Balance
	resp.UserFrozenAmount = user.FrozenAmount + req.FrozenAmount

	err = lockUser(req.UserId)
	if err != nil {
		return
	}

	err = utility.GetUsersDb().Transaction(func(tx *gorm.DB) error {
		err = utility.GetUsersDb().Model(&user).Update("frozen_amount", gorm.Expr("frozen_amount+?", req.FrozenAmount)).Error
		if err != nil {
			err = errors.New("更新frozen_amount失败" + err.Error())
			return err
		}
		var usersDetail = model.UsersDetail{
			UserId:       req.UserId,
			Balance:      resp.UserBalance,
			FrozenAmount: resp.UserFrozenAmount,
			Type:         utility.TccTypeTry,
		}
		err = tx.Create(&usersDetail).Error
		if err != nil {
			fmt.Println("users_detail 插入失败")
		}
		resp.UsersDetailId = int32(usersDetail.ID)
		return nil
	})

	return
}

func (s server) AutoMigrate(ctx context.Context, req *users.Users) (resp *users.UsersSuccess, err error) {
	fmt.Println("users__AutoMigrate")
	resp = &users.UsersSuccess{}
	utility.GetUsersDb().Debug().Set("gorm:table_options", "comment '用户表'").AutoMigrate(&model.Users{})
	fmt.Println("users__AutoMigrate11")
	utility.GetUsersDb().Debug().Set("gorm:table_options", "comment '用户余额流水表'").AutoMigrate(&model.UsersDetail{})
	fmt.Println("users__AutoMigrate33")

	resp.Value = 1
	return
}

func (s server) FrozenAmountCancel(ctx context.Context, req *users.FrozenCancelParams) (resp *users.FrozenCancelResp, err error) {
	resp = &users.FrozenCancelResp{}

	var userDetail = model.UsersDetail{}
	err = utility.GetUsersDb().First(&userDetail, req.UsersDetailId).Error
	if err != nil {
		fmt.Println("查找用户流水失败", err)
		return
	}

	user, err := getUsersInfo(userDetail.UserId)
	if err != nil {
		return
	}

	err = utility.GetUsersDb().Transaction(func(tx *gorm.DB) error {
		err = utility.GetUsersDb().Model(&model.Users{}).Where("id = ?", userDetail.UserId).Update("frozen_amount", gorm.Expr("frozen_amount-?", userDetail.FrozenAmount)).Error
		if err != nil {
			fmt.Println("更新用户冻结金额失败", err)
			return err
		}

		var usersDetail = model.UsersDetail{
			UserId:       userDetail.UserId,
			Balance:      user.Balance,
			FrozenAmount: user.FrozenAmount - userDetail.FrozenAmount,
			Type:         utility.TccTypeCancel,
		}
		err = tx.Create(&usersDetail).Error
		if err != nil {
			fmt.Println("users_detail 插入失败")
			return err
		}
		return nil
	})

	if err != nil {
		resp.Value = 0
		return
	}

	resp.Value = 1
	return
}

func main() {
	listener, err := net.Listen("tcp", utility.UsersGrpcAddr)
	if err != nil {
		log.Println("net listen err ", err)
		return
	}

	s := grpc.NewServer()
	users.RegisterUsersInfoServer(s, &server{})
	log.Printf("start gRPC listen on port %s\n", utility.UsersGrpcAddr)
	if err := s.Serve(listener); err != nil {
		log.Println("failed to serve...", err)
		return
	}
}
