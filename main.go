package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/popeskul/houser/pkg/configs"
	"github.com/popeskul/houser/pkg/middleware"
	"github.com/popeskul/houser/pkg/routes"
	"github.com/popeskul/houser/pkg/utils"
)

// @title Houser API
// @version 1.0
// @description This is an auto-generated API Docs.

// @host localhost:8080
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
func main() {
	// Define env and viper
	configs.EnvConfigs()

	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Routes.
	routes.SwaggerRoute(app)  // Register a route for API Docs (Swagger).
	routes.PublicRoutes(app)  // Register a public routes for app.
	routes.PrivateRoutes(app) // Register a private routes for app.
	routes.NotFoundRoute(app) // Register route for 404 Error.

	// Start server (with graceful shutdown).
	utils.StartServerWithGracefulShutdown(app)
}
