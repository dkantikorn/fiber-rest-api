package models

import "time"

type User struct {
	ID        uint      `json:"id"`
	Firstname string    `json:"first_name"`
	Lastname  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
}
