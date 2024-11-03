package utils

import (
    "log"
    "time"
    "os"
    "github.com/dgrijalva/jwt-go"
    "github.com/joho/godotenv"
)

var jwtSecret []byte

func init() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Erro ao carregar .env")
    }
    jwtSecret = []byte(os.Getenv("JWT_SECRET"))
    if len(jwtSecret) == 0 {
        log.Fatal("JWT_SECRET não foi definido no .env")
    }
}

type Claims struct {
    UserID int    `json:"user_id"`
    Role   string `json:"role"`
    jwt.StandardClaims
}

func GenerateToken(userID int, role string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        UserID: userID,
        Role:   role,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}

func ParseToken(tokenStr string) (*Claims, error) {
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })
    if err != nil || !token.Valid {
        return nil, err
    }
    return claims, nil
}