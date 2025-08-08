package model

import "time"

type User struct {
	ID        string    `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"-" db:"password_hash"`
	PublicKey string    `json:"publicKey" db:"public_key"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}