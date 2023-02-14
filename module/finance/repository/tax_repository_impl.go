package repository

import (
	"database/sql"
	"fmt"

	"github.com/valyala/fasthttp"
)

type taxRepositoryImpl struct {
	DB *sql.DB
}

func NewTaxRepository(db *sql.DB) TaxRepository {
	return &taxRepositoryImpl{
		DB: db,
	}
}

func (repository *taxRepositoryImpl) GetByCoa(ctx *fasthttp.RequestCtx, coa string) (int, error) {
	var rates int
	query := fmt.Sprintf(`
		SELECT t.rates_percent as rates FROM finance.taxes t WHERE t.coa_code = '%s'
	`, coa)

	err := repository.DB.QueryRowContext(ctx, query).Scan(&rates)
	return rates, err
}
