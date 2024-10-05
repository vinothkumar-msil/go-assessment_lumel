package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"go-backend-assessment/config"
	"log"
)

var DB *sql.DB

func Init() {
	config.Load() // Load configuration

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		config.AppConfig.DBUser,
		config.AppConfig.DBPassword,
		config.AppConfig.DBName,
		config.AppConfig.DBHost,
		config.AppConfig.DBPort,
	)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Failed to ping the database:", err)
	}

	log.Println("Connected to the database")
}
