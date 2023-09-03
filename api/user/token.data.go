package user

import (
	"errors"
	"strings"

	postgres "github.com/codefresco/go-build-service/database"
	"github.com/codefresco/go-build-service/loggerfactory"
	"gorm.io/gorm"
)

func CreateToken(token *Token) error {
	logger := loggerfactory.GetSugaredLogger()
	result := postgres.DB.Create(token)

	if result.Error != nil && strings.Contains(result.Error.Error(), "ERROR: duplicate key") {
		return ErrAlreadyEsists
	}

	if result.Error != nil {
		logger.Errorw("Error writing token to database!", "error", result.Error)
		return ErrInternal
	}
	return nil
}

func FindToken(token *Token) (Token, error) {
	logger := loggerfactory.GetSugaredLogger()
	dbToken := Token{}
	result := postgres.DB.First(&dbToken, Token{UserID: token.UserID})

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return dbToken, ErrNotFound
	}

	if result.Error != nil {
		logger.Errorw("Error reading token from database!", "error", result.Error)
		return dbToken, ErrInternal
	}

	return dbToken, nil
}

func UpdateToken(token *Token) error {
	logger := loggerfactory.GetSugaredLogger()
	result := postgres.DB.Save(token)

	if result.Error != nil {
		logger.Errorw("Error updating token in database!", "error", result.Error)
		return ErrInternal
	}

	return nil
}

func DeleteToken(token *Token) error {
	logger := loggerfactory.GetSugaredLogger()
	result := postgres.DB.Delete(&token)

	if result.Error != nil {
		logger.Errorw("Error deleting token from database!", "error", result.Error)
		return ErrInternal
	}

	return nil
}
