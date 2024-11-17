package main

import (
	"fmt"
	"log"
	"nexus-music/db"
	"nexus-music/handlers"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Fetch database configuration from environment variables
	dbUser := getEnv("DB_USER", "admin")
	dbPassword := getEnv("DB_PASSWORD", "password")
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbName := getEnv("DB_NAME", "Nexus_Music")

	// Construct the database connection string
	dataSourceName := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName,
	)

	// Initialize the database connection
	db.InitDB(dataSourceName)

	// Run migrations to ensure tables exist
	db.RunMigrations()

	// Fetch the application port from environment variables
	appPort := getEnv("APP_PORT", "8080")

	// Set up Gin router
	r := gin.Default()

	// Add CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Adjust this to match your frontend's origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Public routes
	r.POST("/login", handlers.Login)
	r.POST("/create-account", handlers.CreateAccount)
	r.GET("/stream/:id", handlers.StreamAudio)

	// Admin routes
	admin := r.Group("/admin")
	admin.POST("/login", handlers.AdminLogin)
	admin.Use(handlers.AuthenticateAdmin())
	{
		admin.GET("/users", handlers.GetAllUsers)
		admin.GET("/streams", handlers.GetAllStreams)
	}

	// Start the server
	log.Printf("Starting server on port %s...", appPort)
	r.Run(":" + appPort)
}

// getEnv retrieves the value of the environment variable or returns a default value if it's not set
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
