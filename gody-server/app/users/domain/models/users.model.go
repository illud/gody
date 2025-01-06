package models

import "time"

type Users struct {
	ID        uint      `json:"id" db:"id" gorm:"primaryKey"`
	Username  string    `json:"username" db:"username"`
	Password  string    `json:"password" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type UsersCreate struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type UsersPut struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type LoginRequest struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}
