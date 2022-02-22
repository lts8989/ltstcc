package model

import "time"

type Users struct {
	ID           int
	Username     string    `gorm:"not null;type:varchar(100);comment:'用户名'"`
	Balance      int32     `gorm:"not null;type:int unsigned default '0' comment '余额，单位：分'"`
	FrozenAmount int32     `gorm:"column:frozen_amount;not null;type:int unsigned default '0' comment '冻结金额，单位分'"`
	Mtime        time.Time `gorm:"default:CURRENT_TIMESTAMP;type:datetime not null on update CURRENT_TIMESTAMP comment '修改时间'"`
	Ctime        time.Time `gorm:"default:CURRENT_TIMESTAMP;type:datetime not null comment '创建时间'"`
}

type UsersDetail struct {
	ID           int
	UserId       int32     `gorm:"not null;column:user_id;type:int unsigned comment '用户id'"`
	Balance      int32     `gorm:"not null;type:int unsigned default '0' comment '余额，单位：分'"`
	FrozenAmount int32     `gorm:"not null;column:frozen_amount;type:int unsigned default '0' comment '冻结金额，单位分'"`
	Type         int8      `gorm:"not null;type:tinyint unsigned comment '类型'"`
	Ctime        time.Time `gorm:"default CURRENT_TIMESTAMP;type:datetime not null comment '创建时间'"`
}
