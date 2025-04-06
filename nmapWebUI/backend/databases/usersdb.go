package databases

import (
	"database/sql"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func VerifyUserCredentials(username, password string) bool {
	var storedHash string

	query := `
		SELECT password_hash FROM users WHERE username = $1
	`

	err := DBObj.QueryRow(query, username).Scan(&storedHash)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("username not found")
			return false
		}
		log.Println(err)
		return false // other DB error
	}

	fmt.Println(storedHash)

	// Compare provided password with stored hash
	if err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password)); err != nil {
		log.Println("invalid password")
		return false
	}

	return true // success
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

func CreateNewUser(username, passwordHash string) (int, error) {
	var userID int
	query := `
        INSERT INTO users (username, password_hash)
        VALUES ($1, $2)
        RETURNING id
    `
	err := DBObj.QueryRow(query, username, passwordHash).Scan(&userID)
	fmt.Println(userID)
	if err != nil {
		log.Println("Error when creating user entry.")
		log.Println(err)
		return 0, err
	}
	return userID, nil
}
