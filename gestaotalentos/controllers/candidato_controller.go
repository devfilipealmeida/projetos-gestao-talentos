package controllers

import (
    "database/sql"
    "encoding/json"
    "gestaotalentos/config"
    "gestaotalentos/models"
    "log"
    "net/http"
    "strconv"
    "strings"

    "github.com/gorilla/mux"
    "github.com/lib/pq"
)

func GetCandidatos(w http.ResponseWriter, r *http.Request) {
    queryParams := r.URL.Query()
    status := queryParams.Get("status")
    search := queryParams.Get("search")

    baseQuery := `SELECT id, nome, habilidades, status FROM candidatos`
    var conditions []string
    var args []interface{}
    argID := 1

    if status != "" {
        conditions = append(conditions, `status = $`+strconv.Itoa(argID))
        args = append(args, status)
        argID++
    }

    if search != "" {
        searchLower := "%" + strings.ToLower(search) + "%"
        conditions = append(conditions, `(LOWER(nome) ILIKE $`+strconv.Itoa(argID)+` OR $`+strconv.Itoa(argID+1)+` = ANY(habilidades))`)
        args = append(args, searchLower, search)
        argID += 2
    }

    if len(conditions) > 0 {
        baseQuery += " WHERE " + strings.Join(conditions, " AND ")
    }

    rows, err := config.DB.Query(baseQuery, args...)
    if err != nil {
        log.Printf("Erro ao buscar candidatos: %v\n", err)
        http.Error(w, "Erro ao buscar candidatos", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var candidatos []models.Candidato
    for rows.Next() {
        var candidato models.Candidato
        if err := rows.Scan(&candidato.ID, &candidato.Nome, pq.Array(&candidato.Habilidades), &candidato.Status); err != nil {
            log.Printf("Erro ao ler candidato: %v\n", err)
            http.Error(w, "Erro ao ler dados de candidatos", http.StatusInternalServerError)
            return
        }
        candidatos = append(candidatos, candidato)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(candidatos)
}


func CreateCandidato(w http.ResponseWriter, r *http.Request) {
    var candidato models.Candidato
    if err := json.NewDecoder(r.Body).Decode(&candidato); err != nil {
        http.Error(w, "Erro ao decodificar dados", http.StatusBadRequest)
        return
    }

    query := `INSERT INTO candidatos (nome, habilidades, status) VALUES ($1, $2, $3) RETURNING id`
    err := config.DB.QueryRow(query, candidato.Nome, pq.Array(candidato.Habilidades), candidato.Status).Scan(&candidato.ID)
    if err != nil {
        log.Printf("Erro ao inserir candidato: %v\n", err)
        http.Error(w, "Erro ao inserir candidato no banco de dados", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(candidato)
}

func GetCandidato(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    var candidato models.Candidato
    query := `SELECT id, nome, habilidades, status FROM candidatos WHERE id = $1`
    err := config.DB.QueryRow(query, id).Scan(&candidato.ID, &candidato.Nome, pq.Array(&candidato.Habilidades), &candidato.Status)
    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "Candidato não encontrado", http.StatusNotFound)
        } else {
            log.Printf("Erro ao buscar candidato: %v\n", err)
            http.Error(w, "Erro ao buscar candidato", http.StatusInternalServerError)
        }
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(candidato)
}

func UpdateCandidato(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    var candidatoAtual models.Candidato
    query := `SELECT id, nome, habilidades, status FROM candidatos WHERE id = $1`
    err := config.DB.QueryRow(query, id).Scan(&candidatoAtual.ID, &candidatoAtual.Nome, pq.Array(&candidatoAtual.Habilidades), &candidatoAtual.Status)
    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "Candidato não encontrado", http.StatusNotFound)
        } else {
            log.Printf("Erro ao buscar candidato: %v\n", err)
            http.Error(w, "Erro ao buscar candidato", http.StatusInternalServerError)
        }
        return
    }

    var candidatoAtualizado models.Candidato
    if err := json.NewDecoder(r.Body).Decode(&candidatoAtualizado); err != nil {
        http.Error(w, "Erro ao decodificar dados", http.StatusBadRequest)
        return
    }

    if candidatoAtualizado.Nome != "" {
        candidatoAtual.Nome = candidatoAtualizado.Nome
    }
    if len(candidatoAtualizado.Habilidades) > 0 {
        candidatoAtual.Habilidades = candidatoAtualizado.Habilidades
    }
    if candidatoAtualizado.Status != "" {
        candidatoAtual.Status = candidatoAtualizado.Status
    }

    query = `UPDATE candidatos SET nome = $1, habilidades = $2, status = $3 WHERE id = $4`
    _, err = config.DB.Exec(query, candidatoAtual.Nome, pq.Array(candidatoAtual.Habilidades), candidatoAtual.Status, id)
    if err != nil {
        log.Printf("Erro ao atualizar candidato: %v\n", err)
        http.Error(w, "Erro ao atualizar candidato no banco de dados", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(candidatoAtual)
}

func DeleteCandidato(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    query := `DELETE FROM candidatos WHERE id = $1`
    result, err := config.DB.Exec(query, id)
    if err != nil {
        log.Printf("Erro ao deletar candidato: %v\n", err)
        http.Error(w, "Erro ao deletar candidato", http.StatusInternalServerError)
        return
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        http.Error(w, "Erro ao verificar exclusão", http.StatusInternalServerError)
        return
    }

    if rowsAffected == 0 {
        http.Error(w, "Candidato não encontrado", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}