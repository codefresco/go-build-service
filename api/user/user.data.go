package user

import (
	"errors"
	"strings"

	postgres "github.com/codefresco/go-build-service/database"
	"github.com/codefresco/go-build-service/libs/pass"
	"github.com/codefresco/go-build-service/loggerfactory"
	"gorm.io/gorm"
)

func CreateUser(user *UserRegisteration) error {
	logger := loggerfactory.GetSugaredLogger()
	dbUser := &User{}
	dbUser.FirstName = user.FirstName
	dbUser.LastName = user.LastName
	dbUser.Email = user.Email
	passwordHash, passwordSalt, err := pass.CreatePassHash(user.Password)

	if err != nil {
		logger.Errorw("Error creating password hash!", "error", err)
		return ErrInternal
	}
	dbUser.PasswordHash = passwordHash
	dbUser.PasswordSalt = passwordSalt
	result := postgres.DB.Create(dbUser)

	if result.Error != nil && strings.Contains(result.Error.Error(), "ERROR: duplicate key") {
		return ErrAlreadyExists
	}

	if result.Error != nil {
		logger.Errorw("Error writing user to database!", "error", result.Error)
		return ErrInternal
	}
	return nil
}

func FindUser(user *UserCredentials) (User, error) {
	logger := loggerfactory.GetSugaredLogger()
	dbUser := User{}
	userDetails := UserDetails{Email: user.Email}
	result := postgres.DB.First(&dbUser, User{UserDetails: userDetails})

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return dbUser, ErrNotFound
	}

	if result.Error != nil {
		logger.Errorw("Error reading user from database!", "error", result.Error)
		return dbUser, ErrInternal
	}

	return dbUser, nil
}
