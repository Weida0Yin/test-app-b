package repo

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"novel-app/internal/domain/entity"
	"novel-app/internal/domain/repository"
	"novel-app/internal/repo/store"
	"novel-app/pkg"
	"strconv"
	"sync"
)

var (
	conn   *gorm.DB
	RdsClt *redis.Client
	once   sync.Once
)

func Init() {
	conn = getConn()
	RdsClt = initRedis()
}

func GetUserRepo() repository.UserRepository {
	return store.NewUserRepo(conn)
}

func getConn() *gorm.DB {
	if conn == nil {
		open()
	}
	return conn
}

func open() {
	once.Do(func() {
		conn = newConn()
	})

}

func initRedis() *redis.Client {
	dbOrder := pkg.GetEnv("REDIS_DB", "")
	num, _ := strconv.Atoi(dbOrder)

	RdsClt = redis.NewClient(&redis.Options{
		Addr:     pkg.GetEnv("REDIS_HOST", ""),
		Password: pkg.GetEnv("REDIS_PASSWORD", ""),
		DB:       num,
	})

	// 测试连接
	_, err := RdsClt.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Connected to Redis")

	return RdsClt
}

func CloseRds() {
	err := RdsClt.Close()
	if err != nil {
		return
	}
}

func newConn() *gorm.DB {
	// get db config
	user := pkg.GetEnv("DB_USER", "")
	pwd := pkg.GetEnv("DB_PASSWORD", "")
	host := pkg.GetEnv("DB_HOST", "")
	port := pkg.GetEnv("DB_PORT", "")
	dbName := pkg.GetEnv("DB_NAME", "")
	//sslMode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pwd, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil
	}

	// 自动迁移，创建数据库表
	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
		return nil
	}
	log.Println("Connected to Mysql")
	return db
}
