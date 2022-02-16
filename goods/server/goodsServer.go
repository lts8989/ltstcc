package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"log"
	"ltstcc/goods"
	"ltstcc/model"
	"ltstcc/utility"
	"net"
)

type server struct {
	goodsMap map[string]*goods.Goods
	goods.UnimplementedGoodsInfoServer
}

func (s server) AutoMigrate(ctx context.Context, req *goods.GoodsId) (resp *goods.GoodsSuccess, err error) {
	fmt.Println("goods__AutoMigrate")
	resp = &goods.GoodsSuccess{}
	utility.GetGoodsDb().Debug().Set("gorm:table_options", "comment '商品信息表'").AutoMigrate(&model.Goods{})
	resp.Value = 1
	return
}

func (s *server) GetGoodsInfo(ctx context.Context, req *goods.GoodsId) (resp *goods.Goods, err error) {
	resp = &goods.Goods{}
	_goods, err := getGoodInfo(req.Value)
	if err != nil {
		return
	}
	resp.Id = int32(_goods.ID)
	resp.GoodsName = _goods.GoodsName
	resp.Price = _goods.Price
	resp.Description = _goods.Description
	resp.FrozenCount = _goods.FrozenCount
	return
}

func getGoodInfo(goodsId int32) (goods model.Goods, err error) {
	goods = model.Goods{}
	err = utility.GetGoodsDb().First(&goods, goodsId).Error
	if err != nil {
		err = errors.New(fmt.Sprintf("查询商品信息失败：%d\n", goodsId) + err.Error())
	}
	return
}
func (s server) GoodsFrozen(ctx context.Context, req *goods.GoodsFrozenParams) (resp *goods.GoodsFrozenResp, err error) {
	resp = &goods.GoodsFrozenResp{Value: 0}
	goodsInfo, err := getGoodInfo(req.GoodsId)
	if err != nil {
		return
	}

	if goodsInfo.StockCount < goodsInfo.FrozenCount+req.Count {
		err = errors.New(fmt.Sprintf("商品库存不足，当前库存：%d,已经冻结数量：%d,冻结数量：%d", goodsInfo.StockCount, goodsInfo.FrozenCount, req.Count))
		return
	}
	err = lockGood(req.GoodsId)
	if err != nil {
		return
	}
	err = utility.GetGoodsDb().Model(&goodsInfo).Update("frozen_count", gorm.Expr("frozen_count+?", req.Count)).Error
	if err != nil {
		err = errors.New("更新frozen_count失败" + err.Error())
		return
	}

	return
}

func main() {
	listener, err := net.Listen("tcp", utility.GoodsGrpcAddr)
	if err != nil {
		log.Println("net listen err ", err)
		return
	}

	s := grpc.NewServer()
	goods.RegisterGoodsInfoServer(s, &server{})
	log.Printf("start gRPC listen on port %s\n", utility.GoodsGrpcAddr)
	if err := s.Serve(listener); err != nil {
		log.Println("failed to serve...", err)
		return
	}
}

func lockGood(goodsId int32) (err error) {
	key := fmt.Sprintf(utility.GOODS_LOCK_KEY, goodsId)
	val, err := utility.GetRedis().Get(key).Result()
	if val != "" {
		err = errors.New("商品已经加锁")
		return
	}
	err = utility.GetRedis().Set(key, utility.GOODS_DETAIL_TYPE_CREATE_ORDER, 0).Err()
	return
}
