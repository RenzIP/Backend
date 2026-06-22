package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var dbInitErr error

func InitDB() {
	if err := godotenv.Load(); err != nil {
		_ = godotenv.Load("../.env")
	}

	dsn := os.Getenv("SUPABASE_URL")
	if dsn == "" {
		dbInitErr = fmt.Errorf("SUPABASE_URL environment variable is not set")
		log.Printf("Warning: %v", dbInitErr)
		return
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		dbInitErr = fmt.Errorf("failed to connect to database: %w", err)
		log.Printf("Warning: %v", dbInitErr)
		return
	}

	DB = db
	dbInitErr = nil
	fmt.Println("Database connection established")
}

func GetDB() *gorm.DB {
	return DB
}

func HasDB() bool {
	return DB != nil
}

func GetDBInitError() error {
	return dbInitErr
}
