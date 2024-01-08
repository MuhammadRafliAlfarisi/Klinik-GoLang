package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"sistem-informasi-klinik/database"
	"sistem-informasi-klinik/router"
	"time"
)

func main() {
	app := fiber.New()

	//Connect to database
	database.ConnectDB()

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,PUT,DELETE",
	})) //Setup the Router
	router.SetupRoutes(app)
	// Listen on PORT 3000
	app.Listen(":3000")

	currentTime := time.Now()

	// Format the current time
	formattedTime := currentTime.Format("2006-01-02 15:04:05")

	// Print the formatted time
	fmt.Println(formattedTime)

}
