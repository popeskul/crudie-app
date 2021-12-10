package models

import (
	"github.com/google/uuid"
	"time"
)

type House struct {
	ID          uuid.UUID `json:"id" db:"id" validate:"required,uuid"`
	Description string    `json:"description" db:"description"`
	Address     string    `json:"address" db:"address"`
	OwnerID     uuid.UUID `json:"owner_id" db:"owner_id" validate:"required,uuid"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type HouseCreateInput struct {
	Description string `json:"description"`
	Address     string `json:"address"`
}

type HouseUpdateInput struct {
	ID          uuid.UUID `json:"id" db:"id" validate:"required,uuid"`
	Description string    `json:"description" db:"description"`
	Address     string    `json:"address" db:"address"`
	OwnerID     uuid.UUID `json:"owner_id" db:"owner_id" validate:"required,uuid"`
}

type HouseDeleteInput struct {
	ID uuid.UUID `json:"id" db:"id" validate:"required,uuid"`
}
