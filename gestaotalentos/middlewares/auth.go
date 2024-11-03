package middlewares

import (
    "context"
    "gestaotalentos/utils"
    "net/http"
    "strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Token não fornecido", http.StatusUnauthorized)
            return
        }

        tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
        claims, err := utils.ParseToken(tokenStr)
        if err != nil {
            http.Error(w, "Token inválido ou expirado", http.StatusUnauthorized)
            return
        }

        r = r.WithContext(context.WithValue(r.Context(), "user", claims))
        next.ServeHTTP(w, r)
    })
}

func AdminMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        claims, ok := r.Context().Value("user").(*utils.Claims)
        if !ok || claims.Role != "admin" {
            http.Error(w, "Acesso não autorizado", http.StatusForbidden)
            return
        }
        next.ServeHTTP(w, r)
    })
}