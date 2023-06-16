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

func (u *UserDao) Insert(data *model.User) error {
	res := u.Db.Create(data)
	return res.Statement.Error
}
func (u *UserDao) GetData(query any, arg ...any) []*model.User {
	res := make([]*model.User, 10)
	u.Db.Model(&model.User{}).Select("password", "salt").Where(query, arg...).Find(&res)
	return res
}
