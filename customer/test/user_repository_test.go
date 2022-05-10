package test

import (
	"testing"
	"thirthfamous/tokopedia-clone-go-graphql/app"
	"thirthfamous/tokopedia-clone-go-graphql/model/domain"
	"thirthfamous/tokopedia-clone-go-graphql/repository"

	"github.com/stretchr/testify/assert"
)

func TestDatabaseConnectionSuccess(t *testing.T) {
	db := app.NewDBTest()
	assert.NoError(t, db.Error)
}

func TestFindByUserNameSuccess(t *testing.T) {
	db := app.NewDBTest()
	db.Exec("SET FOREIGN_KEY_CHECKS = 0;")
	db.Exec("TRUNCATE customer_user;")
	db.Exec("TRUNCATE customer_profile")
	userRepository := repository.NewUserRepository()

	user := domain.User{
		Username: "Farhan",
		Password: "123",
		Profile: domain.Profile{
			Name:    "Hamdallah",
			Address: "Sukabumi",
		},
	}
	db.Create(&user)

	foundUser := domain.User{}
	userRepository.FindByUsername(db, "Farhan", &foundUser)

	assert.Equal(t, user.Username, foundUser.Username)
}

func TestFindByUserNameFailed(t *testing.T) {
	db := app.NewDBTest()
	db.Exec("SET FOREIGN_KEY_CHECKS = 0;")
	db.Exec("TRUNCATE customer_user;")
	db.Exec("TRUNCATE customer_profile")
	userRepository := repository.NewUserRepository()

	user := domain.User{
		Username: "Farhan",
		Password: "123",
		Profile: domain.Profile{
			Name:    "Hamdallah",
			Address: "Sukabumi",
		},
	}
	db.Create(&user)

	foundUser := domain.User{}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	userRepository.FindByUsername(db, "Not Found", &foundUser)
}

func TestCreateUserSuccess(t *testing.T) {
	db := app.NewDBTest()
	db.Exec("SET FOREIGN_KEY_CHECKS = 0;")
	db.Exec("TRUNCATE customer_user;")
	db.Exec("TRUNCATE customer_profile;")
	userRepository := repository.NewUserRepository()
	user := domain.User{
		Username: "Farhan",
		Password: "123",
		Profile: domain.Profile{
			Name:    "Hamdallah",
			Address: "Sukabumi",
		},
	}

	userRepository.CreateUser(db, &user)
	foundUser := domain.User{}
	db.Where("username = ?", user.Username).First(&foundUser)

	assert.Equal(t, foundUser.Username, user.Username)
}
