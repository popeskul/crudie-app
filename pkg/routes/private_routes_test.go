package routes

import (
	"github.com/popeskul/houser/app/models"
	"github.com/popeskul/houser/pkg/utils"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestPrivateRoutes(t *testing.T) {
	// Load .env.test file from the root folder.
	if err := godotenv.Load("../../configs/config.test.yml"); err != nil {
		panic(err)
	}

	// Create a sample data string.
	dataString := `{"id": "00000000-0000-0000-0000-000000000000"}`

	// Create access token.
	token, err := utils.GenerateNewAccessToken(models.User{Email: "test@mail.com", Password: "test@mail.com"})
	if err != nil {
		panic(err)
	}

	// Define a structure for specifying input and output data of a single test case.
	tests := []struct {
		description   string
		route         string // input route
		method        string // input method
		tokenString   string // input token
		body          io.Reader
		expectedError bool
		expectedCode  int
	}{
		{
			description:   "delete house without JWT and body",
			route:         "/api/v1/house",
			method:        "DELETE",
			tokenString:   "",
			body:          nil,
			expectedError: false,
			expectedCode:  fiber.StatusBadRequest,
		},
		{
			description:   "update house without JWT and body",
			route:         "/api/v1/house",
			method:        "PUT",
			tokenString:   "",
			body:          nil,
			expectedError: false,
			expectedCode:  fiber.StatusBadRequest,
		},
		{
			description:   "update house without JWT and body",
			route:         "/api/v1/house",
			method:        "POST",
			tokenString:   "",
			body:          nil,
			expectedError: false,
			expectedCode:  fiber.StatusBadRequest,
		},
		{
			description:   "delete house without right credentials",
			route:         "/api/v1/house",
			method:        "DELETE",
			tokenString:   "Bearer " + token,
			body:          strings.NewReader(dataString),
			expectedError: false,
			expectedCode:  fiber.StatusForbidden,
		},
		{
			description:   "update house without right credentials",
			route:         "/api/v1/house",
			method:        "PUT",
			tokenString:   "Bearer " + token,
			body:          strings.NewReader(dataString),
			expectedError: false,
			expectedCode:  fiber.StatusForbidden,
		},
		{
			description:   "update house without right credentials",
			route:         "/api/v1/house",
			method:        "POST",
			tokenString:   "Bearer " + token,
			body:          strings.NewReader(dataString),
			expectedError: false,
			expectedCode:  fiber.StatusForbidden,
		},
		{
			description:   "delete house without right credentials",
			route:         "/api/v1/house",
			method:        "DELETE",
			tokenString:   "Bearer " + token,
			body:          strings.NewReader(dataString),
			expectedError: false,
			expectedCode:  fiber.StatusForbidden,
		},
	}

	// Define a new Fiber app.
	app := fiber.New()

	// Define routes.
	PrivateRoutes(app)

	// Iterate through test single test cases
	for _, test := range tests {
		// Create a new http request with the route from the test case.
		req := httptest.NewRequest(test.method, test.route, test.body)
		req.Header.Set("Authorization", test.tokenString)
		req.Header.Set("Content-Type", "application/json")

		// Perform the request plain with the app.
		resp, err := app.Test(req, -1) // the -1 disables request latency

		// Verify, that no error occurred, that is not expected
		assert.Equalf(t, test.expectedError, err != nil, test.description)

		// As expected errors lead to broken responses,
		// the next test case needs to be processed.
		if test.expectedError {
			continue
		}

		// Verify, if the status code is as expected.
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
