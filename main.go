package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"io"
	"ltstcc/goods"
	"ltstcc/orders"
	"ltstcc/users"
	"ltstcc/utility"
	"net/http"
	"strconv"
	"time"
)

func init() {
	fmt.Println("main init")
}

type aaa struct {
	Fn func(ctx *gin.Context, intf interface{}) error
}

//go:generate protoc --go_out=./goods/ ./goods/goods.proto
//go:generate protoc --go-grpc_out=./goods/ ./goods/goods.proto
//go:generate protoc --go_out=./orders/ ./orders/orders.proto
//go:generate protoc --go-grpc_out=./orders/ ./orders/orders.proto
//go:generate protoc --go_out=./users/ ./users/users.proto
//go:generate protoc --go-grpc_out=./users/ ./users/users.proto
func main() {
	gin.DisableConsoleColor()
	path := "gin"
	writer, _ := rotatelogs.New(
		path+"%Y%m%d%H.log",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(time.Duration(1800)*time.Second),

		//这里设置1分钟产生一个日志文件
		rotatelogs.WithRotationTime(time.Duration(360)*time.Second),
	)
	gin.DefaultWriter = io.MultiWriter(writer)
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})
	// 创建用户
	r.GET("/createuser", createUser)

	// 批量创建用户
	r.GET("/batchcreateuser", batchcreateuser)

	// 用户充值
	r.GET("/recharge", recharge)

	// 用户下单
	r.GET("/createorder", createOrder)

	r.GET("/automigrate", autoMigrate)

	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8000")
}

func createUser(c *gin.Context) {
	u := &users.Users{Username: "aaaa"}
	userId, err := utility.GetUsersGrpcClient().CreateUsers(c, u)
	if err != nil {
		fmt.Println("add user fail.", err)
		c.String(http.StatusInternalServerError, "add user fail.")
		return
	}
	fmt.Printf("new userId:%d\n", userId.Value)

	c.String(http.StatusOK, "createuser success!")
}

func autoMigrate(c *gin.Context) {
	u := &users.Users{Username: "aaaa"}
	utility.GetUsersGrpcClient().AutoMigrate(c, u)

	gi := &goods.GoodsId{}
	utility.GetGoodsGrpcClient().AutoMigrate(c, gi)

	createOrderParams := &orders.CreateOrderParams{}
	utility.GetOrdersGrpcClient().AutoMigrate(c, createOrderParams)

	c.String(http.StatusOK, "AutoMigrate success!")
}

func batchcreateuser(c *gin.Context) {
	params := &users.BatchCreateUsersParams{Count: 10}
	res, err := utility.GetUsersGrpcClient().BatchCreateUsers(c, params)
	if err != nil {
		fmt.Println("", err)
		c.String(http.StatusInternalServerError, "")
		return
	}
	fmt.Printf("batch create user count:%d", res.Count)
	c.String(http.StatusOK, "batchcreateuser success!")
}

func recharge(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))
	rechargeMoney, _ := strconv.Atoi(c.Query("recharge_money"))

	fmt.Printf("recharge params,userid:%d,rechargemoney:%d", userId, rechargeMoney)

	params := &users.RechargeParams{
		UserId:        int32(userId),
		RechargeMoney: int32(rechargeMoney),
	}
	res, err := utility.GetUsersGrpcClient().Recharge(c, params)
	if err != nil {
		fmt.Println("recharge fail", err)
		c.String(http.StatusInternalServerError, "")
		return
	}
	fmt.Printf("充值成功:%d", res.Value)
	c.String(http.StatusOK, "recharge success!")

}

func createOrder(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))
	goodsId, _ := strconv.Atoi(c.Query("goods_id"))
	goodsCount, _ := strconv.Atoi(c.Query("goods_count"))

	//region 获取商品信息
	goodsInfoParams := &goods.GoodsId{
		Value: int32(goodsId),
	}
	resGoodsInfo, err := utility.GetGoodsGrpcClient().GetGoodsInfo(c, goodsInfoParams)
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	changeAmount := resGoodsInfo.Price * int32(goodsCount)
	//endregion

	//region 冻结用户余额
	frozenParams := &users.FrozenTryParams{
		UserId:       int32(userId),
		FrozenAmount: changeAmount,
	}
	frozenRes, err := utility.GetUsersGrpcClient().FrozenAmountTry(c, frozenParams)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer unlockUser(int32(userId))
	//ch := make(chan aaa)
	//frozenCommitParams := &users.FrozenCommitParams{UsersDetailId: frozenRes.UsersDetailId}
	//ch <- aaa{utility.GetUsersGrpcClient().FrozenAmountCommit(c, frozenCommitParams)}
	//utility.GetUsersGrpcClient().FrozenAmountCancel(c,)
	a := 1
	a = 2
	fmt.Println(a)

	//endregion

	//region 创建订单
	createOrderParams := &orders.CreateOrderParams{
		UserId:           int32(userId),
		GoodsId:          int32(goodsId),
		GoodsCount:       int32(goodsCount),
		ChangeAmount:     changeAmount,
		UserBalance:      frozenRes.UserBalance,
		UserFrozenAmount: frozenRes.UserFrozenAmount,
	}
	createOrderRes, err := utility.GetOrdersGrpcClient().CreateOrder(c, createOrderParams)
	if err != nil {
		fmt.Println("创建订单出错", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	_ = createOrderRes.OrdersDetailId
	//endregion

	//region 冻结商品数量
	goodsFrozenParams := &goods.GoodsFrozenParams{
		GoodsId: int32(goodsId),
		Count:   int32(goodsCount),
	}
	_, err = utility.GetGoodsGrpcClient().GoodsFrozen(c, goodsFrozenParams)
	if err != nil {
		fmt.Println("冻结商品库存数量失败", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer unlockGood(int32(goodsId))
	//endregion

	fmt.Println("下单成功")
	c.String(http.StatusOK, "create order success!")
}

func unlockUser(usersId int32) {
	fmt.Println("unlockUser")
	utility.GetRedis().Del(fmt.Sprintf(utility.USERS_LOCK_KEY, usersId))
	return
}

func unlockGood(goodsId int32) {
	fmt.Println("unlockGood")
	utility.GetRedis().Del(fmt.Sprintf(utility.GOODS_LOCK_KEY, goodsId))
	return
}
