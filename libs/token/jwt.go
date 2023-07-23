package jwt

import (
	"time"

	"github.com/codefresco/go-build-service/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func generateJWT(subject string, ttl time.Duration) (string, error) {
	configs := config.GetConfig()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  subject,
		"exp":  time.Now().Add(ttl).Unix(),
		"iat":  time.Now().Unix(),
		"iss":  configs.Version,
		"uuid": uuid.New(),
	})

	signedToken, err := token.SignedString([]byte(configs.JWTSecret))

	return signedToken, err
}

func GenerateAuthToken(email string) (string, error) {
	configs := config.GetConfig()

	signedToken, err := generateJWT(email, configs.AccessTokenExpiry)

	return signedToken, err
}

func GenerateRefreshToken(email string) (string, error) {
	configs := config.GetConfig()

	signedToken, err := generateJWT(email, configs.RefreshTokenExpiry)

	return signedToken, err
}
