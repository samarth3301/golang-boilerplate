package service

import (
	"database/sql"
	"os"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitPostgres() error {
	var err error
	DB, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}
	
	if err := DB.Ping(); err != nil {
		return err
	}
	
	return nil
}

func ClosePostgres() {
	if DB != nil {
		DB.Close()
	}
} 