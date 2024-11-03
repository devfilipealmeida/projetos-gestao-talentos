package main

import (
    "gestaotalentos/config"
    "gestaotalentos/routes"
    "log"
    "net/http"

    "github.com/joho/godotenv"
    httpSwagger "github.com/swaggo/http-swagger"
    _ "gestaotalentos/docs"
)

// @title API de Gestão de Talentos
// @version 1.0
// @description API para gerenciar candidatos, autenticação e autorização.
// @host localhost:8080
// @BasePath /
func main() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Erro ao carregar o arquivo .env")
    }

    config.ConnectDB()

    router := routes.SetupRoutes()

    corsHandler := func(h http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            w.Header().Set("Access-Control-Allow-Origin", "*")
            w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
            w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

            if r.Method == "OPTIONS" {
                w.WriteHeader(http.StatusOK)
                return
            }

            h.ServeHTTP(w, r)
        })
    }

    router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

    log.Println("API rodando na porta 8080")
    log.Fatal(http.ListenAndServe(":8080", corsHandler(router)))
}