package store

import (
	"gorm.io/gorm"
	"novel-app/internal/domain/entity"
	"novel-app/internal/domain/repository"
	"sync"
)

var (
	ur     *userRepo
	urOnce sync.Once
)

var _ repository.UserRepository = (*userRepo)(nil)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) repository.UserRepository {
	urOnce.Do(func() {
		ur = &userRepo{db: db}
	})
	return ur
}

// Create 注册用户
func (u *userRepo) Create(user *entity.User) error {
	return u.db.Create(user).Error
}

// FindByUName 查找用户名
func (u *userRepo) FindByUName(userName string) (entity.User, error) {
	var user entity.User
	err := u.db.Where("user_name = ?", userName).First(&user).Error
	return user, err
}

// FindById 获取用户信息
func (u *userRepo) FindById(userId int64) (entity.User, error) {
	var user entity.User
	err := u.db.Where("id = ?", userId).First(&user).Error
	return user, err
}

// Save 更新用户信息
func (u *userRepo) Save(user *entity.User) error {
	return u.db.Save(user).Debug().Error
}
