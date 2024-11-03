package models

import (
    "golang.org/x/crypto/bcrypt"
)

type Usuario struct {
    ID    int    `json:"id"`
    Nome  string `json:"nome"`
    Email string `json:"email"`
    Senha string `json:"senha,omitempty"`
    Role  string `json:"role"`
}

func (u *Usuario) HashPassword() error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Senha), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    u.Senha = string(hashedPassword)
    return nil
}

func (u *Usuario) CheckPassword(password string) error {
    return bcrypt.CompareHashAndPassword([]byte(u.Senha), []byte(password))
}