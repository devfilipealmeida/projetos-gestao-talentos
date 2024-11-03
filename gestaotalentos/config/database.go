package config

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/lib/pq"
    "github.com/golang-migrate/migrate/v4"
    "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
)

var DB *sql.DB

func ConnectDB() {
    var err error
    connStr := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
    )
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Erro ao conectar com o banco de dados: ", err)
    }
    if err = DB.Ping(); err != nil {
        log.Fatal("Banco de dados inacessível: ", err)
    }
    fmt.Println("Conectado com sucesso!")

    RunMigrations()
}

func RunMigrations() {
    driver, err := postgres.WithInstance(DB, &postgres.Config{})
    if err != nil {
        log.Fatal("Erro ao configurar o driver de migração: ", err)
    }

    m, err := migrate.NewWithDatabaseInstance(
        "file://migrations",
        "postgres", driver,
    )
    if err != nil {
        log.Fatal("Erro ao iniciar migrações: ", err)
    }
    
    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        log.Fatal("Erro ao aplicar migrações: ", err)
    }
    fmt.Println("Migrações aplicadas com sucesso!")
}