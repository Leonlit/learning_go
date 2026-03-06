package service

import (
	"context"
	"errors"
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

var ErrInvalidCredentials = errors.New("invalid username or password")

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Login(ctx context.Context, username, password string) (string, error) {
	userID, err := s.userRepo.VerifyUserCredentials(ctx, username, password)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	token, err := generateJWT(userID)
	if err != nil {
		return "", err
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
		return &utils.AppError{
			Code:    "PASSWORD_MISMATCH",
			Message: "Passwords do not match",
			Status:  400,
		}
	}

	userExists, err := repository.CheckUsernameExists(username)
	if err != nil {
		return &utils.AppError{
			Code:    "DB_ERROR",
			Message: "Failed to check username",
			Status:  500,
		}
	}

	if userExists {
		return &utils.AppError{
			Code:    "USERNAME_EXISTS",
			Message: "Username already exists",
			Status:  400,
		}
	}

	passwordHash := hashPassword(password)
	if passwordHash == "" {
		return &utils.AppError{
			Code:    "PASSWORD_HASH_FAILED",
			Message: "Failed to hash password",
			Status:  500,
		}
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
