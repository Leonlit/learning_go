package utils

import (
	"log"
	"os"
)

func GetFileBytes(filepath string) []byte {
	file, err := os.ReadFile(filepath)

	if err != nil {
		if os.IsNotExist(err) {
			log.Println("File does not exists: ", err)
			return nil
		} else {
			log.Println("Unable to read file: ", err)
			return nil
		}

	}
	return file
}
