package model

import "time"

type Orders struct {
	ID           int
	State        int8      `gorm:"not null;type:tinyint unsigned comment '状态。1、try。2、commit。3、cancel。'"`
	OrderNo      string    `gorm:"not null;column:order_no;type:varchar(100);comment:'订单号，业务需要'"`
	UserId       int32     `gorm:"not null;column:user_id;type:int unsigned comment '用户id'"`
	GoodsId      int32     `gorm:"not null;column:goods_id;type:int unsigned comment '商品id'"`
	GoodsCount   int32     `gorm:"not null;column:goods_count;type:int unsigned comment '购买商品数量'"`
	ChangeAmount int32     `gorm:"not null;column:change_amount;type:int unsigned comment '订单金额'"`
	Mtime        time.Time `gorm:"default:CURRENT_TIMESTAMP;type:datetime not null on update CURRENT_TIMESTAMP comment '修改时间'"`
	Ctime        time.Time `gorm:"default:CURRENT_TIMESTAMP;type:datetime not null comment '创建时间'"`
}

type OrdersDetail struct {
	ID               uint
	OrdersId         int32     `gorm:"not null;column:orders_id;type:int unsigned comment '订单id'"`
	Type             int8      `gorm:"not null;type:tinyint unsigned comment '操作类型。1、try。2、commit。3、cancel。'"`
	UserId           int32     `gorm:"not null;column:user_id;type:int unsigned comment '用户id'"`
	UserBalance      int32     `gorm:"not null;type:int unsigned default '0' comment '用户余额'"`
	UserFrozenAmount int32     `gorm:"not null;type:int unsigned default '0' comment '用户冻结金额'"`
	Mtime            time.Time `gorm:"default:CURRENT_TIMESTAMP;type:datetime not null on update CURRENT_TIMESTAMP comment '修改时间'"`
	Ctime            time.Time `gorm:"default:CURRENT_TIMESTAMP;type:datetime not null comment '创建时间'"`
}
