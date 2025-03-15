package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	appCommand "proj/internal/user/application/command"
	"proj/internal/user/application/dto"
	appQuery "proj/internal/user/application/query"
	"proj/pkg/logger"
	"proj/pkg/validator"

	domainCommand "proj/internal/user/domain/command"
	domainQuery "proj/internal/user/domain/query"

	"go.uber.org/zap"
)

type UserHandler struct {
	commandHandler *appCommand.UserCommandHandler
	queryHandler   *appQuery.UserQueryHandler
}

func NewUserHandler(cmdHandler *appCommand.UserCommandHandler, queryHandler *appQuery.UserQueryHandler) *UserHandler {
	return &UserHandler{
		commandHandler: cmdHandler,
		queryHandler:   queryHandler,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	logger.Info("Creating new user",
		zap.String("path", r.URL.Path),
		zap.String("method", r.Method),
	)

	var req dto.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("Failed to decode request body",
			zap.Error(err),
			zap.String("path", r.URL.Path),
		)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validator.ValidateStruct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cmd := domainCommand.CreateUserCommand{
		Ctx:      r.Context(),
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	user, err := h.commandHandler.HandleCreate(cmd)
	if err != nil {
		logger.Error("Failed to create user",
			zap.Error(err),
			zap.String("email", req.Email),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info("User created successfully",
		zap.String("email", user.Email),
		zap.String("username", user.Username),
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	q := domainQuery.GetUserByIDQuery{ID: id}
	user, err := h.queryHandler.HandleGetByID(q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/api/users/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	id, err := extractIDFromURL(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var req dto.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validator.ValidateStruct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cmd := domainCommand.UpdateUserCommand{
		ID:       id,
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
	user, err := h.commandHandler.HandleUpdate(cmd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/api/users/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	id, err := extractIDFromURL(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cmd := domainCommand.DeleteUserCommand{ID: id}
	if err := h.commandHandler.HandleDelete(cmd); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	q := domainQuery.ListUsersQuery{
		Limit:  20,
		Offset: 0,
	}

	users, err := h.queryHandler.HandleList(q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) CurrentUser(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("claims")
	fmt.Println(claims)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(claims)
}

func extractIDFromURL(r *http.Request) (string, error) {
	id := r.PathValue("id")
	if id == "" {
		return "", fmt.Errorf("invalid id format")
	}
	return id, nil
}
