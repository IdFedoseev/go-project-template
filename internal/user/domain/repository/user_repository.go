package repository

import (
	"context"
	"proj/internal/user/domain/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	GetByID(id entity.ID) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	Delete(id entity.ID) error
	List() ([]*entity.User, error)
}
