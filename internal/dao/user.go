package dao

import "gindemo/internal/model"

type UserDao struct {
	*Dao
}

func NewUserDao(dao *Dao) *UserDao {
	return &UserDao{
		Dao: dao,
	}
}

func (u *UserDao) Insert() error {
	res := u.db.Create(u)
	return res.Statement.Error
}
func (u *UserDao) GetData(query any, arg ...any) []*model.User {
	res := make([]*model.User, 10)
	u.db.Model(u).Select("password", "salt").Where(query, arg...).Find(&res)
	return res
}
