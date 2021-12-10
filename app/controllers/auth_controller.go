package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/popeskul/houser/app/models"
	"github.com/popeskul/houser/pkg/utils"
	"github.com/popeskul/houser/platform/database"
	"time"
)

// SignIn method for login to the system.
// @Description SignIn to the system with the token.
// @Summary login and creates a new access token
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body models.SignInInput true "user"
// @Success 200 {string} status "ok"
// @Router /v1/sign-in [post]
func SignIn(c *fiber.Ctx) error {
	// Create new User struct
	parsedUser := &models.SignInInput{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(parsedUser); err != nil {
		// Return status 400 and error message
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get user by ID.
	user, err := db.Login(parsedUser.Email, parsedUser.Password)
	if err != nil {
		// Return, if user not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "user with the given email and password is not found",
			"user":  nil,
		})
	}

	// Generate a new Access token.
	token, err := utils.GenerateNewAccessToken(user)
	if err != nil {
		// Return status 500 and token generation error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":        false,
		"msg":          nil,
		"access_token": token,
	})
}

// SignUp method for sign up
// @Description SignUp to the system with the token.
// @Summary signup and creates a new access token
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body models.SignUpInput true "user"
// @Success 200 {string} status "ok"
// @Router /v1/sign-up [post]
func SignUp(c *fiber.Ctx) error {
	// Create new User struct
	user := &models.User{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(user); err != nil {
		// Return status 400 and error message
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create a new validator for a User model.
	validate := utils.NewValidator()

	// Set initialized default data for user:
	user.ID = uuid.New()
	user.CreatedAt = time.Now()

	// Validate user fields.
	if err := validate.Struct(user); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Get user by ID.
	_, err = db.RegisterUser(user)
	if err != nil {
		// Return, if user not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "user with the given email and password is not found",
			"user":  nil,
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"user":  user,
	})
}
