package svc

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"novel-app/internal/domain/entity"
	"novel-app/internal/domain/repository"
	"novel-app/internal/repo"
	"novel-app/pkg/common"
	"time"
)

type UserService struct {
	// mysql
	store repository.UserRepository
	//cache
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		store: repo,
	}
}

// Register 用户注册
func (us *UserService) Register(req *entity.RegisterUserReq) error {
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(hashPwd)
	reqData := entity.User{
		Username: req.Username,
		Password: string(hashPwd),
	}

	// 验证用户名是否注册
	isSaved, _ := us.checkUserName(req.Username)
	if !isSaved {
		return us.store.Create(&reqData)
	}

	return errors.New("user already exists")
}

// 验证用户是否存在
func (us *UserService) checkUserName(username string) (bool, error) {
	_, err := us.store.FindByUName(username)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Login 用户登录
func (us *UserService) Login(ctx context.Context, username, password string) (string, error) {
	user, _ := us.store.FindByUName(username)
	if user.Username != username {
		return "", errors.New("username not match")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("password not match")
	}

	// get token in redis
	rdsKey := fmt.Sprintf("user_auth_%d", user.ID)
	token, err := repo.RdsClt.Get(ctx, rdsKey).Result()
	if err == nil && token != "" {
		return token, nil
	}

	// created token and store token in redis
	token, err = common.GenerateJWT(user.ID)
	if err != nil {
		return "", errors.New("token create fail")
	}
	err = repo.RdsClt.Set(ctx, rdsKey, token, time.Duration(3600)*time.Second).Err()
	if err != nil {
		return "", errors.New("token create fail")
	}

	return token, nil
}

// Logout 退出登录
func (us *UserService) Logout(ctx context.Context, userID int64) error {
	rdsKey := fmt.Sprintf("user_auth_%d", userID)
	token, _ := repo.RdsClt.Get(ctx, rdsKey).Result()
	if token != "" {
		err := repo.RdsClt.Del(ctx, rdsKey).Err()
		if err != nil {
			return errors.New("logout fail")
		}
	}
	return nil
}

// GetInfo 获取用户信息
func (us *UserService) GetInfo(userId int64) (entity.User, error) {
	user, err := us.store.FindById(userId)
	if err != nil {
		return entity.User{}, errors.New("record not found")
	}

	// 用户信息中的密码替换成***
	user.Password = "********"

	return user, nil
}

// ChangePwd 修改密码
func (us *UserService) ChangePwd(userId int64, pwd *entity.ChangePwd) error {
	// 获取用户信息
	user, err := us.store.FindById(userId)
	if err != nil {
		return errors.New("record not found")
	}

	// 验证旧密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pwd.OldPassword)); err != nil {
		return errors.New("password not match")
	}

	// 更新密码
	if pwd.NewPassword == "" {
		return errors.New("new password is empty")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pwd.NewPassword)); err == nil {
		return errors.New("the password is same ")
	}

	hashPwd, err := bcrypt.GenerateFromPassword([]byte(pwd.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	newPassword := string(hashPwd)
	user.Password = newPassword

	// 更新store
	err = us.store.Save(&user)
	if err != nil {
		return errors.New("update password error")
	}

	// 重置登录状态(删除redis登录记录)
	rdsKey := fmt.Sprintf("user_auth_%d", userId)
	repo.RdsClt.Del(context.Background(), rdsKey)

	return nil
}
