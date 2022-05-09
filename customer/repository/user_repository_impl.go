package repository

import (
	"thirthfamous/tokopedia-clone-go-graphql/helper"
	"thirthfamous/tokopedia-clone-go-graphql/model/domain"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) FindByUsername(db *gorm.DB, username string, user *domain.User) {
	db.Where("username = ?", username).First(&user)
}

func (repository *UserRepositoryImpl) CreateUser(db *gorm.DB, user *domain.User) {
	result := db.Create(&user)
	helper.PanicIfError(result.Error)
}
