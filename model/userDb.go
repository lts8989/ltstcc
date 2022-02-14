package model

type Users struct {
	ID           uint
	Username     string
	Balance      int32
	FrozenAmount int32 `gorm:"column:frozen_amount"`
}

type UsersDetail struct {
	ID           uint
	UserId       int32 `gorm:"column:user_id"`
	Balance      int32
	FrozenAmount int32 `gorm:"column:frozen_amount"`
	Type         int8
}
