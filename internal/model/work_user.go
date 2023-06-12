package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"unique;column:username;type:varchar(60);"`
	Password string `gorm:"column:password;type:char(64);"`
	Salt     string `gorm:"column:salt;type:char(64);"`
}

func (u *User) Insert() error {
	res := db.Create(u)
	return res.Statement.Error
}
func (u *User) GetData(query any, arg ...any) []*User {
	res := make([]*User, 10)
	db.Model(u).Select("password", "salt").Where(query, arg...).Find(&res)
	return res
}
func (u *User) TableName() string {
	return "work_user"
}
