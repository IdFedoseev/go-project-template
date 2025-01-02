package router

import (
	"proj/internal/common/interfaces/http/router"
	"proj/internal/user/application/command"
	"proj/internal/user/application/query"
	"proj/internal/user/domain/repository"
	"proj/internal/user/interfaces/http/handler"
)

type UserRouter struct {
	userRepo repository.UserRepository
}

func NewUserRouter(userRepo repository.UserRepository) *UserRouter {
	return &UserRouter{
		userRepo: userRepo,
	}
}

func (r *UserRouter) RegisterRoutes(group *router.Group) {
	// Command and Query handlers
	cmdHandler := command.NewUserCommandHandler(r.userRepo)
	queryHandler := query.NewUserQueryHandler(r.userRepo)

	// Handlers
	userHandler := handler.NewUserHandler(cmdHandler, queryHandler)

	// Routes
	group.HandleFunc("GET /users", userHandler.ListUsers)
	group.HandleFunc("POST /users", userHandler.CreateUser)
	group.HandleFunc("GET /users/{id}", userHandler.GetUser)
	group.HandleFunc("PUT /users/{id}", userHandler.UpdateUser)
	group.HandleFunc("DELETE /users/{id}", userHandler.DeleteUser)
}
