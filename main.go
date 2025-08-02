package main

import (
	"context"
	"crypto/tls"
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
		Addr:     utils.GetEnv("DB_HOST", "localhost:5432"),
		User:     utils.GetEnv("DB_USER", "postgres"),
		Password: utils.GetEnv("DB_PASS", ""),
		Database: utils.GetEnv("DB_NAME", "postgres"),
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true, // For development only; turn off in production
		},
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
	postRepo := repo.NewPostRepo(db)
	commentRepo := repo.NewCommentRepo(db)
	commentService := services.NewCommentService(commentRepo)
	commentHandler := handler.NewCommentHandler(commentService)
	postService := services.NewPostService(postRepo)
	postHandler := handler.NewPostHandler(postService)
	userService := services.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	route := router.SetupRouter(userHandler, postHandler, commentHandler)
	// utils.Test(500, db)
	fmt.Println("Server is running on port", utils.GetEnv("PORT", ":8080"))
	err = http.ListenAndServe(utils.GetEnv("PORT", ":8080"), route)
	if err != nil {
		log.Fatal("Failed to start server", err)
	}
}
