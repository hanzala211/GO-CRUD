package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/hanzala211/CRUD/internal/api/handler"
	"github.com/hanzala211/CRUD/internal/repo"
	"github.com/hanzala211/CRUD/internal/services"
	"github.com/hanzala211/CRUD/router"
	"github.com/hanzala211/CRUD/utils"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: Failed to load .env file: %v", err)
		log.Println("Continuing with system environment variables...")
	}

	fmt.Println("Starting CRUD application...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: utils.GetEnv("DB_PASS", ""),
		Database: utils.GetEnv("DB_NAME", ""),
	})
	defer db.Close()
	if err := db.Ping(ctx); err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	err = utils.CreateSchema(db)
	if err != nil {
		log.Fatal("Failed to create schema", err)
	}

	userRepo := repo.NewUserRepo(db)
	userService := services.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	route := router.SetupRouter(userHandler)
	fmt.Println("Server is running on port", utils.GetEnv("PORT", ":8080"))
	err = http.ListenAndServe(utils.GetEnv("PORT", ":8080"), route)
	if err != nil {
		log.Fatal("Failed to start server", err)
	}
}
