package database

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres driver
)

// Config holds database configuration
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// NewConnection establishes a connection to the database
func NewConnection(config *Config) (*sqlx.DB, error) {
	// Build connection string
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	// Open connection
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Verify connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error verifying database connection: %w", err)
	}

	log.Println("Successfully connected to database!")
	return db, nil
}

// 1. **Connection Configuration**
// - We define a `Config` struct to hold database credentials
// - This makes it easier to pass configuration around
//
// 2. **Connection Setup**
// - `sqlx.Connect` establishes the database connection
// - We're using the postgres driver (imported with `_ "github.com/lib/pq"`)
//
// 3. **Connection Pool Settings**
// - `SetMaxOpenConns(25)`: Maximum number of open connections
// - `SetMaxIdleConns(25)`: Maximum number of idle connections
// - `SetConnMaxLifetime(5 * time.Minute)`: How long a connection can be reused
//
// 4. **Connection Verification**
// - `db.Ping()` ensures we can actually reach the database
