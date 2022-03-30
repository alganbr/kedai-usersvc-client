package models

import (
	"time"
)

type UserPassword struct {
	Id        int64     `json:"id"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
}

type UserPasswordRq struct {
	Id          int64  `json:"id"`
	Password    string `json:"password"`
	RequestedBy string `json:"requested_by"`
}

type ValidatePasswordRq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
