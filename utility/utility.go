package utility

import (
	"fmt"
	"google.golang.org/grpc"
	"ltstcc/goods"
	"ltstcc/orders"
	"ltstcc/users"
	"math/rand"
	"time"
)

const (
	UsersGrpcAddr  = "users_server:8081"
	GoodsGrpcAddr  = "goods_server:8081"
	OrdersGrpcAddr = "orders_server:8081"

	TccTypeTry    = 1
	TccTypeCommit = 2
	TccTypeCancel = 3

	//region 用户流水
	USERS_LOCK_KEY             = "user_lock:%d"
	USERS_DETAIL_TYPE_RECHARGE = 4 //充值
	//endregion

	//region 商品流水
	GOODS_LOCK_KEY                 = "good_lock:%d"
	GOODS_DETAIL_TYPE_CREATE_ORDER = 1 //下单

	//endregion

)

func RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

var usersInfoClient users.UsersInfoClient
var goodsInfoClient goods.GoodsInfoClient
var ordersInfoClient orders.OrdersInfoClient

func GetUsersGrpcClient() users.UsersInfoClient {
	if usersInfoClient == nil {
		conn, err := grpc.Dial(UsersGrpcAddr, grpc.WithInsecure())
		if err != nil {
			fmt.Println("did not connect.", err)
			panic(err)
		}
		usersInfoClient = users.NewUsersInfoClient(conn)
	}

	return usersInfoClient
}

func GetGoodsGrpcClient() goods.GoodsInfoClient {
	if goodsInfoClient == nil {
		conn, err := grpc.Dial(GoodsGrpcAddr, grpc.WithInsecure())
		if err != nil {
			fmt.Println("did not connect.", err)
			panic(err)
		}
		goodsInfoClient = goods.NewGoodsInfoClient(conn)
	}

	return goodsInfoClient
}

func GetOrdersGrpcClient() orders.OrdersInfoClient {
	if ordersInfoClient == nil {
		conn, err := grpc.Dial(OrdersGrpcAddr, grpc.WithInsecure())
		if err != nil {
			fmt.Println("did not connect.", err)
			panic(err)
		}
		ordersInfoClient = orders.NewOrdersInfoClient(conn)
	}
	return ordersInfoClient
}
