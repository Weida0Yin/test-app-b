package svc

import (
	"novel-app/internal/domain/repository"
	"novel-app/internal/repo"
)

var ur repository.UserRepository

func Init() {
	ur = repo.GetUserRepo()
}

func GetUser() *UserService {
	return NewUserService(ur)
}

func GetUpload() *UploadService {
	return NewUploadService()
}
