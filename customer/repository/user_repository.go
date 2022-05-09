package repository

import (
	"thirthfamous/tokopedia-clone-go-graphql/model/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByUsername(db *gorm.DB, username string, user *domain.User)
	CreateUser(db *gorm.DB, user *domain.User)
}
