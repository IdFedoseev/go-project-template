package handler

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"proj/internal/auth/application/dto"
	"proj/internal/auth/domain"
	"proj/internal/user/application/query"
	domainQuery "proj/internal/user/domain/query"
	"proj/pkg/validator"
	"time"
)

type AuthHandler struct {
	UserQueryHandler *query.UserQueryHandler
}

func NewAuthHandler(queryHandler *query.UserQueryHandler) *AuthHandler {
	return &AuthHandler{
		UserQueryHandler: queryHandler,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// Здесь будет логика аутентификации
	var req dto.LoginDto

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := validator.ValidateStruct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	cmd := domainQuery.GetUserByEmailQuery{
		Email: req.Email,
	}
	user, err := h.UserQueryHandler.HandleGetByEmail(cmd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Здесь будет логика проверки пароля
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	// Здесь будет логика генерации токена
	claims := domain.Claims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(1))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "proj",
			Subject:   user.Username,
			ID:        user.ID.(primitive.ObjectID).String(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret")) //TODO secret token
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Здесь будет логика отправки токена
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := map[string]string{
		"token": tokenString,
	}
	json.NewEncoder(w).Encode(resp)
}
