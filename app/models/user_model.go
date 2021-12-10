package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id" db:"id" validate:"required,uuid"`
	Name      string    `json:"name" db:"name" validate:"lte=30"`
	Email     string    `json:"email" db:"email" validate:"required,email"`
	Password  string    `json:"password" db:"password" validate:"required,min=3,max=30"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type UserCreateInput struct {
	Name     string `json:"name" db:"name" validate:"lte=30"`
	Email    string `json:"email" db:"email" validate:"required,email"`
	Password string `json:"password" db:"password" validate:"required,min=3,max=30"`
}

type UserUpdateInput struct {
	ID       uuid.UUID `json:"id" db:"id" validate:"required,uuid"`
	Name     string    `json:"name" db:"name" validate:"lte=30"`
	Email    string    `json:"email" db:"email" validate:"email"`
	Password string    `json:"password" db:"password" validate:"lte=30"`
}

type UserDeleteInput struct {
	ID uuid.UUID `json:"id" db:"id" validate:"required,uuid"`
}
