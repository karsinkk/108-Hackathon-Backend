package dif

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

var (
	db     *sql.DB
	dbOnce sync.Once
)

// GetDB returns a singleton database connection pool
func GetDB() *sql.DB {
	dbOnce.Do(func() {
		config := GetConfig()

		psqlInfo := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			config.Host,
			config.Port,
			config.Username,
			config.Password,
			config.DBName,
		)

		var err error
		db, err = sql.Open("postgres", psqlInfo)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to open database connection")
		}

		// Configure connection pool
		db.SetMaxOpenConns(25)
		db.SetMaxIdleConns(5)
		db.SetConnMaxLifetime(5 * time.Minute)
		db.SetConnMaxIdleTime(2 * time.Minute)

		// Verify connection
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := db.PingContext(ctx); err != nil {
			log.Fatal().Err(err).Msg("Failed to ping database")
		}

		log.Info().
			Str("host", config.Host).
			Str("database", config.DBName).
			Msg("Database connection pool established")
	})

	return db
}

// ConnectDB is kept for backward compatibility but returns the shared connection pool
// Deprecated: Use GetDB() instead
func ConnectDB() *sql.DB {
	return GetDB()
}

// CloseDB closes the database connection pool
func CloseDB() error {
	if db != nil {
		return db.Close()
	}
	return nil
}
