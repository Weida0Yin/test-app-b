package svc

import (
	"fmt"
	"io"
	"novel-app/pkg/oss"
	"time"
)

type UploadService struct{}

func NewUploadService() *UploadService {
	return &UploadService{}
}

// GenerateObjectKey 生成唯一的对象存储路径 ex:{module}/{userID}/{timestampNano}-{filename}
func (s *UploadService) GenerateObjectKey(module string, userID int64, filename string) string {
	return fmt.Sprintf("%s/%d/%d-%s", module, userID, time.Now().UnixNano(), filename)
}

// UploadFile 上传文件到oss
func (s *UploadService) UploadFile(objectKey string, file io.Reader) (string, error) {
	//上传文件流
	err := oss.OSSBucket.PutObject(objectKey, file)
	if err != nil {
		return "", fmt.Errorf("upload file error: %w", err)
	}

	//返回完整的访问url
	return fmt.Sprintf("%s/%s", oss.OSSPrefixUrl, objectKey), nil
}
