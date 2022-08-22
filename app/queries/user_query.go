package queries

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/popeskul/houser/app/models"
)

// UserQueries struct for queries from User model.
type UserQueries struct {
	*sqlx.DB
}

// GetUsers method for getting all users.
func (q *UserQueries) GetUsers() ([]models.User, error) {
	var users []models.User

	query := `SELECT * FROM users`

	err := q.Select(&users, query)
	if err != nil {
		return users, err
	}

	return users, nil
}

// GetUserById method for getting one user by given ID.
func (q *UserQueries) GetUserById(id uuid.UUID) (models.User, error) {
	user := models.User{}

	query := `SELECT * FROM users WHERE id = $1`

	err := q.Get(&user, query, id)
	if err != nil {
		return user, err
	}

	return user, nil
}

// CreateUser method for creating user by given User object.
func (q *UserQueries) CreateUser(b *models.User) error {
	query := `INSERT INTO users VALUES ($1, $2, $3, $4, $5)`

	_, err := q.Exec(query, b.ID, b.Name, b.Email, b.Password, b.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

// UpdateUser method for updating user by given User object.
func (q *UserQueries) UpdateUser(id uuid.UUID, user *models.User) error {
	query := `UPDATE users SET name = $2, email = $3, password = $4 WHERE id = $1`

	_, err := q.Exec(query, id, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUser method for delete user by given ID.
func (q *UserQueries) DeleteUser(id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = $1`

	_, err := q.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
