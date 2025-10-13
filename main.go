package main

import (
    "database/sql"
    _ "github.com/jackc/pgx/v5/stdlib"
    "fmt"
    "log"
)

func main() {
    
    dsn := "postgres://myuser:Udayan10@localhost:5432/mydb?sslmode=disable"

    db, err := sql.Open("pgx", dsn)
    if err != nil {
        log.Fatalf("Failed to open DB: %v", err)
    }
    defer db.Close()

    if err := db.Ping(); err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }

    fmt.Println("✅ Connected to PostgreSQL successfully!")
}
