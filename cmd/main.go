package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"novel-app/internal/repo"
	"novel-app/internal/routers"
	"novel-app/internal/svc"
	"novel-app/pkg"
	"novel-app/pkg/oss"
)

func main() {
	config := pkg.LoadConfig()
	// 初始化各种配置
	initEnv(config)
	defer repo.CloseRds()
	// 注册路由
	r := gin.Default()
	routers.RegisterRouter(r)
	if err := r.Run(":" + config.SYS_PORT); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func initEnv(cfg pkg.SYSConfig) {
	initService(cfg)
}

func initService(cfg pkg.SYSConfig) {
	repo.Init()
	svc.Init()
	oss.InitOss(cfg)
}
