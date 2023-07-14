package user

import (
	"errors"
	"strings"

	postgres "github.com/codefresco/go-build-service/database"
	"github.com/codefresco/go-build-service/libs/pass"
	"github.com/codefresco/go-build-service/loggerfactory"
	"gorm.io/gorm"
)

func CreateUser(user *RegisterUser) error {
	logger := loggerfactory.GetSugaredLogger()
	dbUser := &User{}
	dbUser.FirstName = user.FirstName
	dbUser.LastName = user.LastName
	dbUser.Email = user.Email
	passwordHash, passwordSalt, err := pass.CreatePassHash(user.Password)

	if err != nil {
		logger.Errorw("Error creating password hash!", "error", err)
		return errors.New("Could not create user")
	}
	dbUser.PasswordHash = passwordHash
	dbUser.PasswordSalt = passwordSalt
	result := postgres.DB.Create(dbUser)

	if strings.Contains(result.Error.Error(), "ERROR: duplicate key") {
		return errors.New("User already exists!")
	}

	if result.Error != nil {
		logger.Errorw("Error writing user to database!", "error", result.Error)
		return result.Error
	}
	return nil
}

func FindUser(user *LoginUser) (User, error) {
	logger := loggerfactory.GetSugaredLogger()
	dbUser := User{}
	result := postgres.DB.First(&dbUser, User{Email: user.Email})

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return dbUser, result.Error
	}

	if result.Error != nil {
		logger.Errorw("Error reading user from database!", "error", result.Error)
		return dbUser, errors.New("Could not find user")
	}

	return dbUser, nil
}
