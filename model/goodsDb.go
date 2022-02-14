package model

type Goods struct {
	ID          uint
	GoodsName   string `gorm:"column:goods_name"`
	Price       int32
	Description string
	StockCount  int32 `gorm:"column:stock_count"`
	FrozenCount int32 `gorm:"column:frozen_count"`
}
