package databases

import (
	"database/sql"
	"fmt"
	"gulnManagement/gulnWebUI/utils"
	"log"
	"strconv"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var DBObj *sql.DB

// DBConfig holds the configuration for the database connection
type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func InitDB() {
	dbPort, err := strconv.Atoi(utils.LoadEnv("DB_PORT"))
	if err != nil {
		log.Fatalf("Invalid DB port number: %v", err)
	}

	dbConfig := DBConfig{
		Host:     utils.LoadEnv("DB_HOST"),
		Port:     dbPort,
		User:     utils.LoadEnv("DB_USER"),
		Password: utils.LoadEnv("DB_PASS"),
		DBName:   utils.LoadEnv("DB_NAME"),
		SSLMode:  "disable", // Use "require" for production
	}

	DBObj, err = NewDB(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	//defer DBObj.Close()
	log.Println("Connected to DB")
}

// NewDB initializes and returns a new database connection
func NewDB(cfg DBConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	// Verify the connection is successful
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	log.Println("Successfully connected to the database")
	return db, nil
}
