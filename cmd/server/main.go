package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/kunalsinghdadhwal/stockpulse/internal/models"
	"github.com/kunalsinghdadhwal/stockpulse/internal/utils"
)

const idleTimeout = 3 * time.Minute

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Initialize the database
	db, err := utils.InitDB()

	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	sqlDb, err := db.DB()

	if err != nil {
		log.Fatalf("Error getting database instance: %v", err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Alert{}, &models.Watchlist{})

	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	app := fiber.New(fiber.Config{
		IdleTimeout: idleTimeout,
	})

	go func() {
		if err := app.Listen(":8080"); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	_ = <-c
	fmt.Println("Gracefully Shutting Down.....")
	_ = app.Shutdown()

	fmt.Println("Running Cleaning Tasks.....")

	defer sqlDb.Close()

	fmt.Println("Shutdown Done!!!")
}
