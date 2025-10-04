package main

import (
	"api_app/internal/handlers/user"
	userRepo "api_app/internal/repository/user"
	"api_app/pkg/database"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.NewPostgresConnection(database.ConnectionInfo{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		Username: os.Getenv("POSTGRES_USERNAME"),
		DBName:   os.Getenv("DBNAME"),
		SSLMode:  "disable",
		Password: os.Getenv("PASSWORD"),
	})

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	userRepo := userRepo.NewRepository(db)
	u := user.NewUser(userRepo)

	router := gin.New()

	user.RegisterEndpoints(router, u)

	router.Run(":3000")
}
