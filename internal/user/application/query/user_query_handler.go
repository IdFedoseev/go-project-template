package query

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"proj/internal/user/domain/entity"
	"proj/internal/user/domain/query"
	"proj/internal/user/domain/repository"
)

type UserQueryHandler struct {
	userRepo repository.UserRepository
}

func NewUserQueryHandler(repo repository.UserRepository) *UserQueryHandler {
	return &UserQueryHandler{
		userRepo: repo,
	}
}

func (h *UserQueryHandler) HandleGetByID(q query.GetUserByIDQuery) (*entity.User, error) {
	objectID, err := primitive.ObjectIDFromHex(q.ID)
	if err != nil {
		return nil, err
	}
	return h.userRepo.GetByID(objectID)
}

func (h *UserQueryHandler) HandleGetByEmail(q query.GetUserByEmailQuery) (*entity.User, error) {
	return h.userRepo.GetByEmail(q.Email)
}

func (h *UserQueryHandler) HandleList(q query.ListUsersQuery) ([]*entity.User, error) {
	// Здесь можно добавить логику пагинации
	return h.userRepo.List()
}
