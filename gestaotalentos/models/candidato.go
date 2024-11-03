package models

type Candidato struct {
    ID          int      `json:"id"`
    Nome        string   `json:"nome"`
    Habilidades []string `json:"habilidades"`
    Status      string   `json:"status"`
}
