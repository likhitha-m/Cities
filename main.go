package main

import (
	"fmt"
	"log"
	"os"

	"cities/config"
	"cities/controllers"
	"cities/storage"
	route "cities/v1"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)
// @title weather api documentation
// @version 1.0
// @description This is a api documentation for weather app






// @host localhost:3100
// @BasePath /api/v1

// @securityDefinitions.apikey Authorization
// @in header
// @name Authorization
func main() {

	fmt.Println("inside main function")
	err := godotenv.Load()
	if err != nil {
		err := godotenv.Load("/var/api/cities/.env")
		if err != nil {
			log.Fatalf("Error getting env, not comming through %v", err)
		}
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	fmt.Println("get env", os.Getenv("ENV"))
	if envName := os.Getenv("ENV"); envName == config.Qa || envName == config.Prod {
		// compresses HTTP response
		e.Use(middleware.Gzip())
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	storage.ConnectLogrus() // log file
	storage.MONGO_DB = storage.ConnectMongoDB()

	fmt.Println("storage.mongo", storage.MONGO_DB)

	// Route
	e.GET("/", controllers.HealthCheck)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	v1 := e.Group("/api/v1")
	route.InitializeRoutes(v1)
	e.Logger.Fatal(e.Start(":3100"))
}
