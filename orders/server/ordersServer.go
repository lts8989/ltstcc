package main

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"log"
	"ltstcc/model"
	"ltstcc/orders"
	"ltstcc/utility"
	"net"
)

type server struct {
	ordersMap map[string]*orders.CreateOrderParams
	orders.UnimplementedOrdersInfoServer
}

func (s server) AutoMigrate(ctx context.Context, req *orders.CreateOrderParams) (resp *orders.OrdersSuccess, err error) {
	fmt.Println("orders__AutoMigrate")
	resp = &orders.OrdersSuccess{}
	utility.GetOrdersDb().Debug().Set("gorm:table_options", "comment '订单表'").AutoMigrate(&model.Orders{})
	utility.GetOrdersDb().Debug().Set("gorm:table_options", "comment '订单操作表'").AutoMigrate(&model.OrdersDetail{})

	resp.Value = 1
	return
}

func (s *server) CreateOrderCancel(ctx context.Context, req *orders.CreateOrderCancelParams) (resp *orders.OrdersSuccess, err error) {
	resp = &orders.OrdersSuccess{}

	var orderDetail = model.OrdersDetail{}
	err = utility.GetOrdersDb().First(&orderDetail, req.OrdersDetailId).Error
	if err != nil {
		fmt.Println("查找订单流水失败", err)
		return
	}

	err = utility.GetOrdersDb().Transaction(func(tx *gorm.DB) error {
		err = tx.Model(&model.Orders{}).Where("id = ?", orderDetail.OrdersId).Update("state", utility.TccTypeCancel).Error
		if err != nil {
			fmt.Println("更新订单状态失败", err)
			return err
		}

		var ordersDetail = model.OrdersDetail{
			OrdersId:         orderDetail.OrdersId,
			Type:             utility.TccTypeCancel,
			UserId:           orderDetail.UserId,
			UserBalance:      req.UserBalance,
			UserFrozenAmount: req.UserFrozenAmount,
		}
		err = tx.Create(&ordersDetail).Error
		if err != nil {
			fmt.Println("order detail 插入失败")
			return err
		}

		return nil
	})

	return
}

func (s *server) CreateOrder(ctx context.Context, req *orders.CreateOrderParams) (resp *orders.CreateOrderResp, err error) {
	resp = &orders.CreateOrderResp{}
	err = utility.GetOrdersDb().Transaction(func(tx *gorm.DB) error {
		var o = model.Orders{
			State:        utility.TccTypeTry,
			OrderNo:      utility.RandString(8),
			UserId:       req.UserId,
			GoodsId:      req.GoodsId,
			GoodsCount:   req.GoodsCount,
			ChangeAmount: req.ChangeAmount,
		}

		err = tx.Create(&o).Error
		if err != nil {
			fmt.Println("order 插入失败")
			return err
		}
		resp.OrdersId = int32(o.ID)

		var ordersDetail = model.OrdersDetail{
			OrdersId:         int32(o.ID),
			Type:             utility.TccTypeTry,
			UserId:           req.UserId,
			UserBalance:      req.UserBalance,
			UserFrozenAmount: req.UserFrozenAmount,
		}

		err = tx.Create(&ordersDetail).Error
		if err != nil {
			fmt.Println("order detail 插入失败")
			return err
		}
		resp.OrdersDetailId = int32(ordersDetail.ID)

		return nil
	})

	return
}

func main() {
	listener, err := net.Listen("tcp", utility.OrdersGrpcAddr)
	if err != nil {
		log.Println("net listen err ", err)
		return
	}

	s := grpc.NewServer()
	orders.RegisterOrdersInfoServer(s, &server{})
	log.Printf("start gRPC listen on port %s\n", utility.OrdersGrpcAddr)
	if err := s.Serve(listener); err != nil {
		log.Println("failed to serve...", err)
		return
	}
}
