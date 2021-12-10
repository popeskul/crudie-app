package utils

import (
	"github.com/popeskul/houser/app/models"
	"github.com/spf13/viper"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

// GenerateNewAccessToken func for generate a new Access token.
func GenerateNewAccessToken(user models.User) (string, error) {
	// Set secret key from .yml file.
	secret := viper.GetString("jwt_secret_key")

	// Set expires minutes count for secret key from .yml file.
	minutesCount, _ := strconv.Atoi(viper.GetString("jwt_secret_key_expire_minutes_count"))

	// Create a new claims.
	claims := jwt.MapClaims{}

	// Set public claims:
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()
	claims["user_id"] = user.ID

	// Create a new JWT access token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate token.
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		// Return error, it JWT token generation failed.
		return "", err
	}

	return t, nil
}
