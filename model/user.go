package model

import "time"

type User struct {
	Id         int64     `json:"id,string" gorm:"primaryKey;autoIncrement" form:"id"`
	UserId     int64     `json:"user_id,string" gorm:"type:bigint(20);not null;unique" form:"user_id"`
	Gender     int8      `json:"gender" gorm:"type:tinyint(4);not null;default:0" form:"gender"`
	UserName   string    `json:"user_name" gorm:"type:varchar(64);not null" form:"user_name"`
	Password   string    `json:"password" gorm:"type:varchar(64);not null" form:"password"`
	Phone      string    `json:"phone" gorm:"type:varchar(11);unique;not null" form:"phone"`
	CreateTime time.Time `json:"create_time" gorm:"not null" form:"create_time"`
	UpdateTime time.Time `json:"update_time" gorm:"not null" form:"update_time"`
}

func (User) TableName() string {
	return "user"
}
