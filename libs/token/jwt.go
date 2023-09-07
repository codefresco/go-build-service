package jwt

import (
	"errors"
	"time"

	"github.com/codefresco/go-build-service/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type MapClaims = jwt.MapClaims

func generateJWT(subject string, ttl time.Duration) (string, uuid.UUID, error) {
	configs := config.GetConfig()
	jti := uuid.New()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MapClaims{
		"sub": subject,
		"exp": time.Now().Add(ttl).Unix(),
		"iat": time.Now().Unix(),
		"iss": configs.Version,
		"jti": jti,
	})

	signedToken, err := token.SignedString([]byte(configs.JWTSecret))

	return signedToken, jti, err
}

func GenerateTokenPair(email string) (string, string, uuid.UUID, uuid.UUID, error) {
	configs := config.GetConfig()

	refreshToken, refreshJti, refreshTokenError := generateJWT(email, configs.RefreshTokenExpiry)
	if refreshTokenError != nil {
		return "", "", uuid.Nil, uuid.Nil, refreshTokenError
	}

	accessToken, accessJti, accessTokenError := generateJWT(email, configs.AccessTokenExpiry)
	if accessTokenError != nil {
		return "", "", uuid.Nil, uuid.Nil, accessTokenError
	}

	return accessToken, refreshToken, accessJti, refreshJti, nil
}

func ValidateToken(jwtToken string) (MapClaims, error) {
	configs := config.GetConfig()

	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, hmac := token.Method.(*jwt.SigningMethodHMAC); !hmac {
			return nil, ErrUnauthorized
		}
		return []byte(configs.JWTSecret), nil
	})

	if errors.Is(err, jwt.ErrTokenExpired) {
		return nil, ErrTokenExpired
	}

	if err != nil {
		return nil, ErrUnauthorized
	}

	if claims, ok := token.Claims.(MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrUnauthorized
}
