package middleware

import "net/http"

func Auth() func(http.HandlerFunc) http.HandlerFunc {
    return func(next http.HandlerFunc) http.HandlerFunc {
        return func(w http.ResponseWriter, r *http.Request) {
            // Здесь будет логика аутентификации
            next(w, r)
        }
    }
} 