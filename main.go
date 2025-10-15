package main

import (
	"database/sql"
	"example/api-transaction/config"
	"example/api-transaction/pkg/repository"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {

	cfg := config.LoadConfig()

	transactionsDB := repository.NewTransactionsDB(initDB(cfg.TRANSACTIONS_DB_USERNAME, cfg.TRANSACTIONS_DB_PASSWORD, cfg.POSTGRES_HOST, cfg.POSTGRES_PORT, cfg.TRANSACTIONS_DB_NAME))

}

func initDB(dbUserName, dbPassword, postgresHost, postgrePort, dbName string) *sql.DB {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUserName, dbPassword, postgresHost, postgrePort, dbName)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	fmt.Println("✅ Connected to PostgreSQL successfully!")
	return db
}
