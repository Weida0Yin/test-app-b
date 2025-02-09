package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"novel-app/internal/domain/entity"
	"novel-app/internal/svc"
	"novel-app/pkg/response"
)

type UserHandler struct {
	service *svc.UserService
}

func NewUserHandler(service *svc.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// Register 新用户注册
func (uh *UserHandler) Register(c *gin.Context) {
	req := &entity.RegisterUserReq{}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, "Invalid input")
		return
	}

	err := uh.service.Register(req)
	if err != nil {
		response.Fail(c, err.Error())
		return
	}
	response.Success(c, nil, "User registered successfully")
}

// Login 用户登录
func (uh *UserHandler) Login(c *gin.Context) {
	req := &entity.LoginUserReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, "Invalid input")
		return
	}
	token, err := uh.service.Login(context.Background(), req.Username, req.Password)
	if err != nil {
		response.Fail(c, err.Error())
		return
	}

	response.Success(c, token, "User login successfully")
	return
}

// Logout 退出登录
func (uh *UserHandler) Logout(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		response.Fail(c, "Unauthorized")
		return
	}

	if userId, ok := userId.(int64); ok {
		_ = uh.service.Logout(context.Background(), userId)
		response.Success(c, nil, "success")
		return
	}

	response.Fail(c, "user not found")
	return
}

// GetUserInfo 获取用户信息
func (uh *UserHandler) GetUserInfo(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		response.Fail(c, "Unauthorized")
		return
	}

	if userId, ok := userId.(int64); ok {
		user, err := uh.service.GetInfo(userId)
		if err != nil {
			response.Fail(c, "user not founded")
			return
		}
		response.Success(c, user, "User info successfully")
		return
	}

	response.Fail(c, "user not found")
	return
}

// UpdateUser 更新用户信息
func (uh *UserHandler) UpdateUser(c *gin.Context) {}

// ChangePwd 更新密码
func (uh *UserHandler) ChangePwd(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		response.Fail(c, "Unauthorized")
		return
	}

	if userId, ok := userId.(int64); ok {
		req := &entity.ChangePwd{}

		if err := c.ShouldBindJSON(&req); err != nil {
			response.Fail(c, "Invalid input")
			return
		}

		err := uh.service.ChangePwd(userId, req)
		if err != nil {
			response.Fail(c, err.Error())
			return
		}
		response.Success(c, nil, "edit password successfully")
		return
	}

	response.Fail(c, "user not found")
	return
}
