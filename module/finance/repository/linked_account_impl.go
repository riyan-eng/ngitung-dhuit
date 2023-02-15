package repository

import (
	"database/sql"
	"fmt"

	"github.com/valyala/fasthttp"
)

type linkedAccountRepositoryImpl struct {
	Database *sql.DB
}

func NewLinkedAccountRepository(database *sql.DB) LinkedAccountRepository {
	return &linkedAccountRepositoryImpl{
		Database: database,
	}
}

func (repository *linkedAccountRepositoryImpl) FindOneByCode(ctx *fasthttp.RequestCtx, code string) (string, error) {
	var coa string
	query := fmt.Sprintf(`
		SELECT la.coa_code as coa FROM finance.linked_accounts la where la.code = '%s'
	`, code)
	err := repository.Database.QueryRowContext(ctx, query).Scan(&coa)
	return coa, err
}
