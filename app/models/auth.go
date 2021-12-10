package models

type SignInInput struct {
	Email    string `json:"email" db:"email" validate:"email"`
	Password string `json:"password" db:"password" validate:"lte=30"`
}

type SignUpInput struct {
	Name     string `json:"name" db:"name" validate:"lte=30"`
	Email    string `json:"email" db:"email" validate:"email"`
	Password string `json:"password" db:"password" validate:"lte=30"`
}
