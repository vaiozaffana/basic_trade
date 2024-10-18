package main

import (
	"BasicTrade/database"
	"BasicTrade/routes"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")

	cloudinaryApiKey := os.Getenv("CLOUDINARY_API_KEY")
	cloudinaryApiSecret := os.Getenv("CLOUDINARY_API_SECRET")
	cloudinaryCloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")

	fmt.Println("Database Host: ", dbHost)
	fmt.Println("Database User: ", dbUser)
	fmt.Println("Database Password: ", dbPassword)
	fmt.Println("Database Name: ", dbName)
	fmt.Println("Database Port: ", dbPort)
	fmt.Println("JWT Secret Key: ", jwtSecretKey)
	fmt.Println("Cloudinary API Key: ", cloudinaryApiKey)
	fmt.Println("Cloudinary API Secret: ", cloudinaryApiSecret)
	fmt.Println("Cloudinary Cloud Name: ", cloudinaryCloudName)

	database.InitDB()

	r := routes.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
