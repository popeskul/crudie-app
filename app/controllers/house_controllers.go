package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/popeskul/houser/app/models"
	"github.com/popeskul/houser/pkg/utils"
	"github.com/popeskul/houser/platform/database"
	"time"
)

// GetHouse func gets house by given ID or 404 error.
// @Description Get house by given ID.
// @Summary get house by given ID
// @Tags House
// @Accept json
// @Produce json
// @Param id path string true "House ID"
// @Success 200 {object} models.House
// @Router /v1/house/{id} [get]
func GetHouse(c *fiber.Ctx) error {
	// Catch user ID from URL.
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
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
	user, err := db.GetHouseById(id)
	if err != nil {
		// Return, if user not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "house with the given ID is not found",
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

// GetHouses godoc.
// @Description Get all exists houses.
// @Summary gets all exists houses
// @Tags Houses
// @Accept json
// @Produce json
// @Success 200 {array} models.House
// @Router /v1/houses [get]
func GetHouses(c *fiber.Ctx) error {
	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get all houses.
	houses, err := db.GetHouses()
	if err != nil {
		// Return 404, if users not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":  true,
			"msg":    "users were not found",
			"count":  0,
			"houses": nil,
		})
	}

	// Return status 200 OK.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":  false,
		"msg":    nil,
		"count":  len(houses),
		"houses": houses,
	})
}

// CreateHouse func for creates a new house.
// @Description Create a new house.
// @Summary creates a new house
// @Tags House
// @Accept json
// @Produce json
// @Param input body models.HouseCreateInput true "house info"
// @Success 200 {object} models.House
// @Security ApiKeyAuth
// @Router /v1/house [post]
func CreateHouse(c *fiber.Ctx) error {
	// Get now time.
	now := time.Now().Unix()

	// Get tokenMetadata from JWT.
	tokenMetadata, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 400 and JWT parse error.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Checking, if now time greater than expiration from JWT.
	if now > tokenMetadata.Expires {
		// Return status 401 and unauthorized error message.
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	// Create new House struct
	house := &models.House{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(house); err != nil {
		// Return status 400 and error message
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 403 and database connection error.
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create a new validator for a House model.
	validate := utils.NewValidator()

	// Set initialized default data for house:
	house.ID = uuid.New()
	house.OwnerID = tokenMetadata.UserId
	house.CreatedAt = time.Now()

	// Validate house fields.
	if err := validate.Struct(house); err != nil {
		// Return 400, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// CreateHouse house.
	if err := db.CreateHouse(house); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 200 OK.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"house": house,
	})
}

// UpdateHouse func for updates house by given ID.
// @Description Update house.
// @Summary update house
// @Tags House
// @Accept json
// @Produce json
// @Param input body models.HouseUpdateInput true "house info"
// @Success 201 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/house [put]
func UpdateHouse(c *fiber.Ctx) error {
	// Get now time.
	now := time.Now().Unix()

	// Get tokenMetadata from JWT.
	tokenMetadata, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Checking, if now time greater than expiration from JWT.
	if now > tokenMetadata.Expires {
		// Return status 401 and unauthorized error message.
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	// Create new House struct
	house := &models.House{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(house); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Checking, if house with given ID is exists.
	foundedHouse, err := db.GetHouseById(house.ID)
	if err != nil {
		// Return status 404 and house not found error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "house with this ID not found",
		})
	}

	// user is house owner
	if tokenMetadata.UserId != foundedHouse.OwnerID {
		// Return status 403 and unauthorized error message.
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": true,
			"msg":   "You don't have permission for update",
		})
	}

	// Create a new validator for a House model.
	validate := utils.NewValidator()

	// Validate user fields.
	if err := validate.Struct(house); err != nil {
		// Return 400, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Checking, if house with given ID is exists.
	err = db.UpdateHouseById(foundedHouse.ID, house)
	if err != nil {
		// Return status 404 and house not found error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "house with this ID not found",
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

// DeleteHouse func for deletes house by given ID.
// @Description Delete house by given ID.
// @Summary delete house by given ID
// @Tags House
// @Accept json
// @Produce json
// @Param input body models.HouseDeleteInput true "house id"
// @Success 204 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/house [delete]
func DeleteHouse(c *fiber.Ctx) error {
	// Get now time.
	now := time.Now().Unix()

	// Get tokenMetadata from JWT.
	tokenMetadata, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Checking, if now time greater than expiration from JWT.
	if now > tokenMetadata.Expires {
		// Return status 401 and unauthorized error message.
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	// Create new User struct
	house := &models.House{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(house); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create a new validator for a User model.
	validate := utils.NewValidator()

	// Validate only one user field ID.
	if err := validate.StructPartial(house, "id"); err != nil {
		// Return 400, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 403 and database connection error.
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Checking, if house with given ID is exists.
	foundedHouse, err := db.GetHouseById(house.ID)
	if err != nil {
		// Return status 404 and user not found error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "user with this ID not found",
		})
	}

	// user is house owner
	if tokenMetadata.UserId != foundedHouse.OwnerID {
		// Return status 403 and unauthorized error message.
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": true,
			"msg":   "You don't have permission for update",
		})
	}

	// Delete house by given ID.
	if err := db.DeleteHouseByID(foundedHouse.ID); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 204 no content.
	return c.SendStatus(fiber.StatusNoContent)
}
