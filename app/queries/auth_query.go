package queries

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/popeskul/houser/app/models"
)

// AuthQueries struct for queries from User model.
type AuthQueries struct {
	*sqlx.DB
}

// Login method for getting one user by given email and password.
func (q *AuthQueries) Login(email, password string) (models.User, error) {
	user := models.User{}

	query := `SELECT id FROM users WHERE email = $1 AND password = $2`

	err := q.Get(&user, query, email, password)
	if err != nil {
		return user, err
	}

	return user, nil
}

// RegisterUser method for creating user by given User object.
func (q *AuthQueries) RegisterUser(b *models.User) (*uuid.UUID, error) {
	var id *uuid.UUID

	query := `INSERT INTO users VALUES ($1, $2, $3, $4, $5) RETURNING id`

	row := q.QueryRow(query, b.ID, b.Name, b.Email, b.Password, b.CreatedAt)

	err := row.Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("tried create user with an error %w", err)
	}

	return id, nil
}
