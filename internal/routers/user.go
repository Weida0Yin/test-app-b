package routers

import (
	"github.com/gin-gonic/gin"
	"novel-app/api/upload/v1"
	"novel-app/api/user/v1"
	"novel-app/internal/svc"
	"novel-app/pkg/middleware"
)

func RegisterRouter(r *gin.Engine) {
	userHandler := user.NewUserHandler(svc.GetUser())
	uploadHandler := upload.NewUploadHandler(svc.GetUpload())

	apiUserV1 := r.Group("/api/user/v1")
	apiUserV1.POST("register", userHandler.Register)
	apiUserV1.POST("login", userHandler.Login)
	// need login
	apiUserV1.POST("getUserInfo", middleware.JWTAuth(), userHandler.GetUserInfo)
	apiUserV1.POST("updateUser", middleware.JWTAuth(), userHandler.UpdateUser)
	apiUserV1.POST("editPwd", middleware.JWTAuth(), userHandler.ChangePwd)
	apiUserV1.POST("logout", middleware.JWTAuth(), userHandler.Logout)

	apiUploadV1 := r.Group("/api/upload/v1")
	apiUploadV1.Use(middleware.JWTAuth())
	apiUploadV1.POST("uploadFile", uploadHandler.UploadFile)
}
