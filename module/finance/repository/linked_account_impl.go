package repository

import (
	"database/sql"
	"fmt"

	"github.com/valyala/fasthttp"
)

type linkedAccountRepositoryImpl struct {
	DB *sql.DB
}

func NewLinkedAccountRepository(db *sql.DB) LinkedAccountRepository {
	return &linkedAccountRepositoryImpl{
		DB: db,
	}
}

func (repository *linkedAccountRepositoryImpl) GetByCode(ctx *fasthttp.RequestCtx, code string) (string, error) {
	var coa string
	query := fmt.Sprintf(`
		SELECT la.coa_code as coa FROM finance.linked_accounts la where la.code = '%s'
	`, code)
	err := repository.DB.QueryRowContext(ctx, query).Scan(&coa)
	return coa, err
}
