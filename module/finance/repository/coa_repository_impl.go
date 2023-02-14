package repository

import (
	"database/sql"
	"fmt"

	"github.com/valyala/fasthttp"
)

func NewCOARepository(db *sql.DB) COARepository {
	return &chartOfAccountRepositoryImpl{
		Database: db,
	}
}

type chartOfAccountRepositoryImpl struct {
	Database *sql.DB
}

func (repository *chartOfAccountRepositoryImpl) GetByCode(ctx *fasthttp.RequestCtx, coa string) error {
	query := fmt.Sprintf(`
		select coa.id from finance.coas coa where coa.code = '%v'
	`, coa)
	err := repository.Database.QueryRowContext(ctx, query).Err()
	if err != nil {
		return err
	}
	return nil
}
