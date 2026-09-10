package model

import "github.com/google/uuid"

type User struct {
	ID           uuid.UUID `json:"id"    db:"id"`
	Login        string    `json:"login" db:"login"`
	PasswordHash string    `json:"-"     db:"password_hash"` // "-" чтобы не отдавать в JSON
}

type SignUpRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type SignInRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
