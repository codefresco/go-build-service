package auth

import (
	"github.com/codefresco/go-build-service/api/token"
	"github.com/codefresco/go-build-service/api/user"
	"github.com/codefresco/go-build-service/libs/pass"
	jwt "github.com/codefresco/go-build-service/libs/token"
	"github.com/codefresco/go-build-service/loggerfactory"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Register(c *fiber.Ctx) error {
	userDetails := c.Locals("body").(*user.UserRegisteration)

	err := user.CreateUser(userDetails)
	if err != nil {
		return authErrorHandler(c, err)
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "User registered!",
		"user":    userDetails.Email,
	})
}

func Login(c *fiber.Ctx) error {
	logger := loggerfactory.GetSugaredLogger()
	loginDetails := c.Locals("body").(*user.UserCredentials)

	dbUser, err := user.FindUser(loginDetails)
	if err != nil {
		return authErrorHandler(c, err)
	}

	validPassword, loginError := pass.PasswordChecks(loginDetails.Password, dbUser.PasswordHash, dbUser.PasswordSalt)
	if loginError != nil {
		logger.Errorw("Error checking password!", "error", loginError)
		return authErrorHandler(c, ErrInternal)
	}

	if !validPassword {
		return authErrorHandler(c, ErrPermissionDenied)
	}

	accessToken, refreshToken, accessJti, refreshJti, tokenError := jwt.GenerateTokenPair(dbUser.Email)
	if tokenError != nil {
		return authErrorHandler(c, ErrPermissionDenied)
	}

	token.CreateToken(&token.Token{UserID: dbUser.ID, AccessJwtID: accessJti, RefreshJwtID: refreshJti})

	return c.Status(200).JSON(fiber.Map{
		"message":       "Login successful!",
		"user":          loginDetails.Email,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func Logout(c *fiber.Ctx) error {
	claims := c.Locals("claims").(jwt.MapClaims)

	userDetails, err := user.FindUser(&user.UserCredentials{Email: claims["sub"].(string)})
	if err != nil {
		return authErrorHandler(c, err)
	}

	accessJwtID, err := uuid.Parse(claims["jti"].(string))
	if err != nil {
		return authErrorHandler(c, err)
	}

	err = token.DeleteToken(&token.Token{UserID: userDetails.ID, AccessJwtID: accessJwtID})
	if err != nil {
		return authErrorHandler(c, err)
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Logout successful!",
		"user":    userDetails.Email,
	})
}

func Refresh(c *fiber.Ctx) error {
	refreshHeader := c.Get("Refresh-Token")

	claims, err := jwt.ValidateToken(refreshHeader)
	if err != nil {
		return c.Status(403).JSON(fiber.Map{
			"message": "Invalid refresh token!",
			"error":   err.Error(),
		})
	}

	userDetails, err := user.FindUser(&user.UserCredentials{Email: claims["sub"].(string)})
	if err != nil {
		return authErrorHandler(c, err)
	}

	refreshJwtID, err := uuid.Parse(claims["jti"].(string))
	if err != nil {
		return authErrorHandler(c, err)
	}

	tokenMeta, err := token.FindToken(&token.Token{UserID: userDetails.ID, RefreshJwtID: refreshJwtID})
	if err != nil {
		return authErrorHandler(c, err)
	}

	accessToken, refreshToken, accessJti, refreshJti, tokenError := jwt.GenerateTokenPair(userDetails.Email)
	if tokenError != nil {
		return authErrorHandler(c, ErrPermissionDenied)
	}

	tokenMeta.AccessJwtID = accessJti
	tokenMeta.RefreshJwtID = refreshJti
	tokenMeta.UserID = userDetails.ID

	err = token.UpdateToken(&tokenMeta)
	if err != nil {
		return authErrorHandler(c, err)
	}

	return c.Status(200).JSON(fiber.Map{
		"message":       "Token refreshed!",
		"user":          userDetails.Email,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func authErrorHandler(c *fiber.Ctx, err error) error {
	switch err {
	case user.ErrAlreadyExists:
		return c.Status(409).JSON(fiber.Map{
			"message": "Could not create user!",
			"error":   err.Error(),
		})
	case user.ErrNotFound:
		return c.Status(404).JSON(fiber.Map{
			"message": "User does not exist!",
			"error":   err.Error(),
		})
	case token.ErrNotFound:
		fallthrough
	case ErrPermissionDenied:
		return c.Status(403).JSON(fiber.Map{
			"message": "Invalid authentication details!",
			"error":   err.Error(),
		})
	default:
		return c.Status(500).JSON(fiber.Map{
			"message": "Something went wrong!",
			"error":   ErrInternal.Error(),
		})
	}
}
