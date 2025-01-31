package api

import (
	"be-go-car-rental/config"
	"be-go-car-rental/docs"
	"be-go-car-rental/routes"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	App *gin.Engine
)

func init() {
	App = gin.New()

	environment := GetEnvOrDefault("ENVIRONMENT", "development")
	if environment == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Println("Warning: Error loading .env file, proceeding without it")
		}
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	mode := os.Getenv("GIN_MODE")
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		log.Println("[INFO] Running in debug mode")
	}

	docs.SwaggerInfo.Title = "Car Rental API"
	docs.SwaggerInfo.Description = "This is API Car Rental."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = GetEnvOrDefault("HOST", "localhost:8080")
	if environment == "development" {
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
	} else {
		docs.SwaggerInfo.Schemes = []string{"https"}
	}

	db := config.ConnectDatabase()

	routes.SetupRouter(db, App)
}

func GetEnvOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func Handler(w http.ResponseWriter, r *http.Request) {
	App.ServeHTTP(w, r)
}
