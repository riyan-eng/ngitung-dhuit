package repository

import (
	"database/sql"
	"fmt"

	"github.com/valyala/fasthttp"
)

type transactionRepositoryImpl struct {
	Database *sql.DB
}

func NewTransactionRepository(database *sql.DB) TransactionRepository {
	return &transactionRepositoryImpl{
		Database: database,
	}
}

func (repository *transactionRepositoryImpl) InsertOne(ctx *fasthttp.RequestCtx, desc string, amount float64) (string, error) {
	queryTransactions := fmt.Sprintf(`
		INSERT INTO finance.transactions (description, amount) VALUES ('%s', '%f') RETURNING id
	`, desc, amount)

	var id string
	err := repository.Database.QueryRowContext(ctx, queryTransactions).Scan(&id)
	return id, err
}
