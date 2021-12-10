package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/popeskul/houser/app/controllers"
	"github.com/popeskul/houser/pkg/middleware"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for /user:
	route.Post("/user", middleware.JWTProtected(), controllers.CreateUser)   // create a new user
	route.Put("/user", middleware.JWTProtected(), controllers.UpdateUser)    // update one user by ID
	route.Delete("/user", middleware.JWTProtected(), controllers.DeleteUser) // delete one user by ID

	// Routes for /house:
	route.Post("/house", middleware.JWTProtected(), controllers.CreateHouse)   // create a new house
	route.Put("/house", middleware.JWTProtected(), controllers.UpdateHouse)    // update a house
	route.Delete("/house", middleware.JWTProtected(), controllers.DeleteHouse) // delete one house by ID
}
