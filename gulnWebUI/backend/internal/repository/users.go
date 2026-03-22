package repository

import (
	"context"
	"database/sql"
	"fmt"
	"gulnManagement/gulnWebUI/internal/utils"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	db *sql.DB
}

var JWTSecretKey = utils.LoadEnv("JWT_SECRET_KEY")

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) VerifyUserCredentials(ctx context.Context, username, password string) (string, error) {
	var storedHash string
	var userUUID string

	query := `
		SELECT uuid, password_hash FROM users WHERE username = $1
	`
	err := r.db.QueryRowContext(ctx, query, username).Scan(&userUUID, &storedHash)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("username not found")
			return "", err
		}
		log.Println(err)
		return "", err // other DB error
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password)); err != nil {
		log.Println("invalid password")
		return "", err
	}

	return userUUID, nil
}

func CheckUsernameExists(username string) (bool, error) {
	// Prepare a SQL query to check if the username exists
	query := "SELECT COUNT(*) FROM users WHERE username = $1"

	// Execute the query and scan the result into a variable
	var count int
	err := DBObj.QueryRow(query, username).Scan(&count)
	if err != nil {
		// Return false and the error if the query failed
		if err == sql.ErrNoRows {
			return false, err // No rows means the username doesn't exist
		}
		log.Println(err)
		return false, err
	}

	// If count > 0, the username exists
	return count > 0, err
}

func GetUserUUID(username string) (string, error) {
	// Prepare a SQL query to check if the username exists
	query := "SELECT uuid FROM users WHERE username = $1"

	// Execute the query and scan the result into a variable
	var uuid string
	err := DBObj.QueryRow(query, username).Scan(&uuid)
	if err != nil {
		// Return false and the error if the query failed
		if err == sql.ErrNoRows {
			return "", err // No rows means the username doesn't exist
		}
		log.Println(err)
		return "", err
	}
	return uuid, err
}

func CreateNewUser(username, passwordHash string) (string, error) {
	var userUUID string
	fmt.Println(username, passwordHash)
	query := `
        INSERT INTO users (username, password_hash, uuid)
        VALUES ($1, $2, uuid_generate_v4())
        RETURNING uuid
    `
	err := DBObj.QueryRow(query, username, passwordHash).Scan(&userUUID)
	if err != nil {
		log.Println("Error when creating user entry.")
		log.Println(err)
		return "", err
	}
	return userUUID, nil
}
