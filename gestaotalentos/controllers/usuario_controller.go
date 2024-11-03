package controllers

import (
    "encoding/json"
    "gestaotalentos/config"
    "gestaotalentos/models"
    "gestaotalentos/utils"
    "log"
    "net/http"
    "github.com/lib/pq"
)

func Register(w http.ResponseWriter, r *http.Request) {
    var usuario models.Usuario
    if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
        http.Error(w, "Erro ao decodificar requisição", http.StatusBadRequest)
        return
    }

    if usuario.Nome == "" || usuario.Email == "" || usuario.Senha == "" || usuario.Role == "" {
        http.Error(w, "Todos os campos são obrigatórios", http.StatusBadRequest)
        return
    }

    if err := usuario.HashPassword(); err != nil {
        http.Error(w, "Erro ao criptografar senha", http.StatusInternalServerError)
        return
    }

    query := `INSERT INTO usuarios (nome, email, senha, role) VALUES ($1, $2, $3, $4) RETURNING id`
    err := config.DB.QueryRow(query, usuario.Nome, usuario.Email, usuario.Senha, usuario.Role).Scan(&usuario.ID)
    if err != nil {
        if err, ok := err.(*pq.Error); ok && err.Code.Name() == "unique_violation" {
            http.Error(w, "Email já cadastrado", http.StatusConflict)
            return
        }
        log.Printf("Erro ao inserir usuário: %v\n", err)
        http.Error(w, "Erro ao salvar usuário no banco de dados", http.StatusInternalServerError)
        return
    }

    usuario.Senha = ""
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(usuario)
}

func Login(w http.ResponseWriter, r *http.Request) {
    var credenciais struct {
        Email string `json:"email"`
        Senha string `json:"senha"`
    }
    if err := json.NewDecoder(r.Body).Decode(&credenciais); err != nil {
        http.Error(w, "Erro ao decodificar requisição", http.StatusBadRequest)
        return
    }

    if credenciais.Email == "" || credenciais.Senha == "" {
        http.Error(w, "Email e senha são obrigatórios", http.StatusBadRequest)
        return
    }

    var usuario models.Usuario

    query := `SELECT id, nome, email, senha, role FROM usuarios WHERE email = $1`
    err := config.DB.QueryRow(query, credenciais.Email).Scan(&usuario.ID, &usuario.Nome, &usuario.Email, &usuario.Senha, &usuario.Role)
    if err != nil {
        http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
        return
    }


    if err := usuario.CheckPassword(credenciais.Senha); err != nil {
        http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
        return
    }

    token, err := utils.GenerateToken(usuario.ID, usuario.Role)
    if err != nil {
        http.Error(w, "Erro ao gerar token", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"token": token})
}