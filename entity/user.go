package entity

import "time"

type User struct {
	ID           int32
	Name         string
	Occupation   string
	Role         string
	Avatar       string
	PasswordHash string
	Email        string
	Visible      string
	Token        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
