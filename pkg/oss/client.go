package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
	"novel-app/pkg"
)

var OSSClient *oss.Client
var OSSBucket *oss.Bucket
var OSSPrefixUrl string

func InitOss(cfg pkg.SYSConfig) {
	//endpoint := pkg.GetEnv("OSS_ENDPOINT", "")
	//ak := pkg.GetEnv("OSS_ACCESS_KEY", "")
	//sk := pkg.GetEnv("OSS_SECRET_KEY", "")
	//bucketName := pkg.GetEnv("OSS_BUCKET", "")
	//prefixUrl := pkg.GetEnv("OSS_PREFIX_URL", "")

	NewClient(cfg.OSS_ENDPOINT, cfg.OSS_ACCESS_KEY, cfg.OSS_SECRET_KEY, cfg.OSS_BUCKET, cfg.OSS_PREFIX_URL)

}

func NewClient(endpoint, ak, sk, bucketName, prefixURL string) {
	//创建oss客户端实例
	ossClient, err := oss.New(endpoint, ak, sk)
	if err != nil {
		log.Fatalf("create oss client error: %v", err)
	}

	// 获取存储空间
	bucket, err := ossClient.Bucket(bucketName)
	if err != nil {
		log.Fatalf("get oss bucket error: %v", err)
	}

	OSSClient = ossClient
	OSSBucket = bucket
	OSSPrefixUrl = prefixURL
}
