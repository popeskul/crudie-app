package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/popeskul/houser/app/controllers"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Test handler
	a.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("App running")
	})

	// Routes auth:
	route.Post("/sign-in", controllers.SignIn) // login to the system
	route.Post("/sign-up", controllers.SignUp) // registration

	// Routes users:
	route.Get("/users", controllers.GetUsers)   // get list of all users
	route.Get("/user/:id", controllers.GetUser) // get one user by ID

	// Routes houses:
	route.Get("/houses", controllers.GetHouses)   // get list of all users
	route.Get("/house/:id", controllers.GetHouse) // get list of all users
}
