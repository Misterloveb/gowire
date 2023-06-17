package model

import "gorm.io/gorm"

type DemoUser struct {
	gorm.Model
	UserName string `gorm:"unique;column:username;type:varchar(60);"`
	Password string `gorm:"column:password;type:char(64);"`
	Salt     string `gorm:"column:salt;type:char(64);"`
}

func (u *DemoUser) TableName() string {
	return "demo_user"
}
