package repository

import "database/sql"

type TransactionsDB struct {
	DB *sql.DB
}

type TransactionRepository interface {
}

func NewTransactionsDB(db *sql.DB) *TransactionsDB {
	return &TransactionsDB{DB: db}

}
