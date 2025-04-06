package databases

import (
	"database/sql"
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

	// Compare provided password with stored hash
	if err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password)); err != nil {
		log.Println("invalid password")
		return false
	}

	return true // success
}

func CreateNewUser(username, passwordHash string) (int, error) {
	var userID int
	query := `
        INSERT INTO users (username, password_hash)
        VALUES ($1, $2)
        RETURNING id
    `
	err := DBObj.QueryRow(query, username, passwordHash).Scan(&userID)
	if err != nil {
		log.Println("Error when creating user entry.")
		log.Println(err)
		return 0, err
	}
	return userID, nil
}
