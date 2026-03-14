package service

import (
	"context"
	"gulnManagement/gulnWebUI/internal/auth"
	"gulnManagement/gulnWebUI/internal/repository"
	"gulnManagement/gulnWebUI/internal/utils"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Login(ctx context.Context, username, password string) (string, *utils.AppError) {
	userID, err := s.userRepo.VerifyUserCredentials(ctx, username, password)
	if err != nil {
		return "", utils.InternalError("Error When Logging In", err)
	}

	token, err := generateJWT(userID)
	if err != nil {
		return "", utils.InternalError("Error When Logging In", err)
	}

	return token, nil
}

func generateJWT(uuid string) (string, error) {
	claims := auth.Claims{
		UserUUID: uuid,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "guln-management",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(8 * time.Hour)), // Expiry set to 8 hour
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(repository.JWTSecretKey))
}

func hashPassword(password string) string {
	// Generate a hashed password with bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("failed to hash password: %v", err)
		return ""
	}
	return string(hashedPassword)
}

func (s *AuthService) Register(ctx context.Context, username, password, repeatPassword string) error {

	if repeatPassword != password {
		return utils.NewAppError(
			"PASSWORD_MISMATCH",
			"Passwords do not match",
			400,
			nil,
		)
	}

	userExists, err := repository.CheckUsernameExists(username)
	if err != nil {
		return utils.NewAppError(
			"DB_ERROR",
			"Failed to check username",
			500,
			nil,
		)
	}

	if userExists {
		return utils.NewAppError(
			"USERNAME_EXISTS",
			"Username already exists",
			400,
			nil,
		)
	}

	passwordHash := hashPassword(password)
	if passwordHash == "" {
		return utils.NewAppError(
			"PASSWORD_HASH_FAILED",
			"Failed to hash password",
			500,
			nil,
		)
	}

	created, err := repository.CreateNewUser(username, passwordHash)
	if err != nil {
		return &utils.AppError{
			Code:    "USER_CREATE_FAILED",
			Message: "Failed to create user",
			Status:  500,
		}
	}

	if created == "" {
		return &utils.AppError{
			Code:    "USER_CREATE_FAILED",
			Message: "User was not created",
			Status:  500,
		}
	}

	return nil
}
