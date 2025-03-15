package router

import (
	"proj/internal/auth/interfaces/http/handler"
	"proj/internal/common/interfaces/http/router"
	"proj/internal/user/application/query"
	"proj/internal/user/domain/repository"
)

type AuthRouter struct {
	userRepo repository.UserRepository
}

func NewAuthRouter(userRepo repository.UserRepository) *AuthRouter {
	return &AuthRouter{
		userRepo: userRepo,
	}
}

func (r *AuthRouter) RegisterRoutes(group *router.Group) {
	queryHandler := query.NewUserQueryHandler(r.userRepo)
	authHandler := handler.NewAuthHandler(queryHandler)

	group.HandleFunc("POST /auth/login", authHandler.Login)
}
