package mongodb

import (
	"proj/internal/user/domain/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MongoUser struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Username  string             `bson:"username"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

// ToEntity конвертирует MongoUser в entity.User
func (u *MongoUser) ToEntity() *entity.User {
	return &entity.User{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

// FromEntity конвертирует entity.User в MongoUser
func FromEntity(user *entity.User) *MongoUser {
	id, _ := user.ID.(primitive.ObjectID)
	return &MongoUser{
		ID:        id,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
