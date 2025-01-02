package entity

import (
	"time"
)

type ID interface{}

type User struct {
	ID        ID        `json:"id" bson:"_id,omitempty"`
	Username  string    `json:"username" bson:"username,omitempty"`
	Email     string    `json:"email" bson:"email,omitempty"`
	Password  string    `json:"-" bson:"password,omitempty"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
