package model

type Orders struct {
	ID           uint
	State        int8
	OrderNo      string
	UserId       int32 `gorm:"column:user_id"`
	GoodsId      int32 `gorm:"column:goods_id"`
	GoodsCount   int32 `gorm:"column:goods_count"`
	ChangeAmount int32 `gorm:"column:change_amount"`
}

type OrdersDetail struct {
	ID               uint
	OrdersId         int32 `gorm:"column:orders_id"`
	Type             int8
	UserId           int32 `gorm:"column:user_id"`
	UserBalance      int32 `gorm:"column:user_balance"`
	UserFrozenAmount int32 `gorm:"column:user_frozen_amount"`
}
