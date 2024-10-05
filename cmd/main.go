package main

import (
	"log"
	"os"

	"go-backend-assessment/api"
	"go-backend-assessment/config"
	"go-backend-assessment/db"
	"go-backend-assessment/scripts"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

func DataRefreshTask() {
	log.Println("Starting data refresh...")
	// Call the CSV loading function or refresh logic here
	err := scripts.LoadCSVData("data.csv") // Replace with your actual file path
	if err != nil {
		log.Printf("Error during data refresh: %v", err)
	} else {
		log.Println("Data refresh completed successfully.")
	}
}

func main() {
	// Initialize configurations
	config.Load()

	// Initialize database connection
	db.Init()

	// Set up the Gin router
	router := gin.Default()
	api.SetupRoutes(router)

	// Schedule a periodic data refresh using cron (e.g., every day at midnight)
	c := cron.New()
	c.AddFunc("*/10 * * * *", func() {
		log.Println("Running scheduled data refresh...")
		go DataRefreshTask() // Run the task in a separate goroutine to prevent blocking
	})

	// Start the cron scheduler
	c.Start()

	// Create a log file for logging periodic data refresh activities
	logFile, err := os.OpenFile("data_refresh.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(logFile)
	} else {
		log.Println("Failed to log to file, using default stderr")
	}

	// Start the HTTP server
	log.Println("Starting server on port 8000")
	log.Fatal(router.Run(":8000"))
}
