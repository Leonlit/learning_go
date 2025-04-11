package databases

import (
	"database/sql"
	"fmt"
	"log"
)

func GetScanList(userID, page string) bool {
	var storedHash string

	query := `
		SELECT password_hash FROM users WHERE username = $1
	`

	err := DBObj.QueryRow(query, userID, page).Scan(&storedHash)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("username not found")
			return false
		}
		log.Println(err)
		return false // other DB error
	}

	fmt.Println(storedHash)

	return true // success
}
