package repository

import (
	"database/sql"
	"fmt"

	"github.com/valyala/fasthttp"
)

type taxRepositoryImpl struct {
	Database *sql.DB
}

func NewTaxRepository(database *sql.DB) TaxRepository {
	return &taxRepositoryImpl{
		Database: database,
	}
}

func (repository *taxRepositoryImpl) FindOneByCoa(ctx *fasthttp.RequestCtx, coa string) (int, error) {
	var rates int
	query := fmt.Sprintf(`
		SELECT t.rates_percent as rates FROM finance.taxes t WHERE t.coa_code = '%s'
	`, coa)

	err := repository.Database.QueryRowContext(ctx, query).Scan(&rates)
	return rates, err
}
