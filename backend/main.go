package main

import (
	"github.com/gofiberAPI-biolerplate/pkg/configs"
	"github.com/gofiberAPI-biolerplate/pkg/middleware"
	"github.com/gofiberAPI-biolerplate/pkg/routes"
	"github.com/gofiberAPI-biolerplate/platform/database"

	"github.com/gofiber/fiber/v2"

	_ "github.com/gofiberAPI-biolerplate/docs"
)


// @title Fiber API Biolerplate
// @version 1.0
// @description API Documentation
// @contact.name API Support
// @contact.email talhajaved700@gmail.com
// @host go.local:5000
// @BasePath /api/v1
func main() {
	
	// Check DB connection & Migrate
	db := database.Connection()
	database.InitialMigration(db)

	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	middleware.FiberMiddleware(app) 

	// Routes.
	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)
	routes.SwaggerRoute(app)

	port := configs.ListenPort()
	app.Listen(port)
}
