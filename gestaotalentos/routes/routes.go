package routes

import (
    "gestaotalentos/controllers"
    "gestaotalentos/middlewares"
    "net/http"

    "github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
    router := mux.NewRouter()

    // Rotas de autenticação
    // @Summary Registrar novo usuário
    // @Description Cadastra um novo usuário no sistema
    // @Tags Autenticação
    // @Accept json
    // @Produce json
    // @Param user body models.User true "Dados do usuário"
    // @Success 200 {object} models.User
    // @Router /register [post]
    router.HandleFunc("/register", controllers.Register).Methods("POST")

    // @Summary Login do usuário
    // @Description Realiza o login do usuário no sistema
    // @Tags Autenticação
    // @Accept json
    // @Produce json
    // @Param credentials body models.Credentials true "Credenciais de login"
    // @Success 200 {object} models.User
    // @Router /login [post]
    router.HandleFunc("/login", controllers.Login).Methods("POST")

    // Rotas para gerenciamento de candidatos
    // @Summary Lista todos os candidatos
    // @Description Retorna uma lista de todos os candidatos
    // @Tags Candidatos
    // @Produce json
    // @Success 200 {array} models.Candidato
    // @Router /candidatos [get]
    router.Handle("/candidatos", middlewares.AuthMiddleware(http.HandlerFunc(controllers.GetCandidatos))).Methods("GET")

    // @Summary Cria um novo candidato
    // @Description Adiciona um novo candidato ao sistema
    // @Tags Candidatos
    // @Accept json
    // @Produce json
    // @Param candidato body models.Candidato true "Dados do candidato"
    // @Success 201 {object} models.Candidato
    // @Router /candidatos [post]
    router.Handle("/candidatos", middlewares.AuthMiddleware(http.HandlerFunc(controllers.CreateCandidato))).Methods("POST")

    // @Summary Busca um candidato pelo ID
    // @Description Retorna os detalhes de um candidato específico
    // @Tags Candidatos
    // @Produce json
    // @Param id path int true "ID do candidato"
    // @Success 200 {object} models.Candidato
    // @Router /candidatos/{id} [get]
    router.Handle("/candidatos/{id}", middlewares.AuthMiddleware(http.HandlerFunc(controllers.GetCandidato))).Methods("GET")

    // @Summary Atualiza um candidato pelo ID
    // @Description Atualiza as informações de um candidato específico
    // @Tags Candidatos
    // @Accept json
    // @Produce json
    // @Param id path int true "ID do candidato"
    // @Param candidato body models.Candidato true "Dados do candidato"
    // @Success 200 {object} models.Candidato
    // @Router /candidatos/{id} [put]
    router.Handle("/candidatos/{id}", middlewares.AuthMiddleware(http.HandlerFunc(controllers.UpdateCandidato))).Methods("PUT")

    // @Summary Deleta um candidato pelo ID
    // @Description Remove um candidato específico do sistema
    // @Tags Candidatos
    // @Param id path int true "ID do candidato"
    // @Success 204
    // @Router /candidatos/{id} [delete]
    router.Handle("/candidatos/{id}", middlewares.AuthMiddleware(middlewares.AdminMiddleware(http.HandlerFunc(controllers.DeleteCandidato)))).Methods("DELETE")

    return router
}