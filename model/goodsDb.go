package model

import "time"

type Goods struct {
	ID          int
	GoodsName   string    `gorm:"not null;column:goods_name;type:varchar(100);comment:'商品名称'"`
	Price       int32     `gorm:"not null;type:int unsigned comment '价格（分）'"`
	Description string    `gorm:"not null;type:varchar(100) default '' comment '商品名称'"`
	StockCount  int32     `gorm:"not null;column:stock_count;type:int unsigned default '0' comment '库存数量'"`
	FrozenCount int32     `gorm:"not null;column:frozen_count;type:int unsigned default '0' comment '冻结数量'"`
	Mtime       time.Time `gorm:"default:CURRENT_TIMESTAMP;type:datetime not null on update CURRENT_TIMESTAMP comment '修改时间'"`
	Ctime       time.Time `gorm:"default:CURRENT_TIMESTAMP;type:datetime not null comment '创建时间'"`
}
