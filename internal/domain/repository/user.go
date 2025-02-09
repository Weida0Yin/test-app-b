package repository

import (
	"novel-app/internal/domain/entity"
)

type UserRepository interface {
	Create(req *entity.User) error                    //用户注册
	FindByUName(userName string) (entity.User, error) //通过用户名获取用户信息
	FindById(userId int64) (entity.User, error)       // 通过ID获取用户信息
	Save(req *entity.User) error                      // 更新用户信息
}
