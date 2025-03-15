package middleware

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"proj/internal/auth/domain"
)

func verifyToken(tokenString string) (*domain.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &domain.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil //TODO secret token
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*domain.Claims); ok && token.Valid {
		fmt.Println(claims)
		return claims, nil
	} else {
		return nil, err
	}
}

func Auth() func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// Здесь будет логика аутентификации
			if r.RequestURI == "/api/auth/login" {
				next(w, r)
				return
			}
			claims, err := verifyToken(r.Header.Get("Authorization"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			next(w, r.WithContext(context.WithValue(r.Context(), "claims", claims)))
		}
	}
}
