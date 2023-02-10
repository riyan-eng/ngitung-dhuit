package repository

import (
	"database/sql"
	"fmt"
)

type transactionRepositoryImpl struct {
	DB *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionRepositoryImpl{
		DB: db,
	}
}

func (repository *transactionRepositoryImpl) Insert(desc string, amount float64) (string, error) {
	queryTransactions := fmt.Sprintf(`
		INSERT INTO finance.transactions (description, amount) VALUES ('%s', '%f') RETURNING id
	`, desc, amount)

	var id string
	err := repository.DB.QueryRow(queryTransactions).Scan(&id)
	return id, err
}
