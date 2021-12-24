package main

import (
	"github.com/Jeanpigi/go-api/src/database"
	"github.com/Jeanpigi/go-api/src/loadenv"
	"github.com/Jeanpigi/go-api/src/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load  Environment Variables
	loadenv.LoadEnv()

	// Get connection to Data Base
	db := database.GetConnection()
	app := fiber.New()

	routes.SetRoutes(app)

	app.Listen(":3000")
	defer db.Close()
}

