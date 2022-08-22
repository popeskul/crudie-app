package queries

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/popeskul/houser/app/models"
)

// HouseQueries struct for queries from User model.
type HouseQueries struct {
	*sqlx.DB
}

// CreateHouse method for creating user by given User object.
func (q *HouseQueries) CreateHouse(h *models.House) error {
	query := `INSERT INTO houses VALUES ($1, $2, $3, $4, $5)`

	_, err := q.Exec(query, h.ID, h.Description, h.Address, h.OwnerID, h.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

// GetHouses method for getting all users.
func (q *HouseQueries) GetHouses() ([]models.House, error) {
	var houses []models.House

	query := `SELECT * FROM houses`

	err := q.Select(&houses, query)
	if err != nil {
		return houses, err
	}

	return houses, nil
}

// GetHouseById method for getting one user by given ID.
func (q *UserQueries) GetHouseById(id uuid.UUID) (models.House, error) {
	house := models.House{}

	query := `SELECT * FROM houses WHERE id = $1`

	err := q.Get(&house, query, id)
	if err != nil {
		return house, err
	}

	return house, nil
}

// UpdateHouseById method for updating house by given House object.
func (q *HouseQueries) UpdateHouseById(id uuid.UUID, house *models.House) error {
	query := `UPDATE houses SET description = $2, address = $3 WHERE id = $1`
	fmt.Println(id, house)

	_, err := q.Exec(query, id, house.Description, house.Address)
	if err != nil {
		return err
	}

	return nil
}

// DeleteHouseByID method for delete user by given ID.
func (q *UserQueries) DeleteHouseByID(id uuid.UUID) error {
	query := `DELETE FROM houses WHERE id = $1`

	_, err := q.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
