package dao

import "github.com/Misterloveb/gowire/internal/model"

type UserDao struct {
	*Dao
}

func NewUserDao(dao *Dao) *UserDao {
	return &UserDao{
		Dao: dao,
	}
}

func (u *UserDao) Insert(data *model.DemoUser) error {
	res := u.Db.Create(data)
	return res.Statement.Error
}
func (u *UserDao) GetData(query any, arg ...any) []*model.DemoUser {
	res := make([]*model.DemoUser, 10)
	u.Db.Model(&model.DemoUser{}).Select("password", "salt").Where(query, arg...).Find(&res)
	return res
}
