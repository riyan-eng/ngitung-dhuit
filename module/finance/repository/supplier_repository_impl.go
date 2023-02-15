package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/valyala/fasthttp"
)

type supplierRepositoryImpl struct {
	Database *sql.DB
}

func NewSupplierRepository(database *sql.DB) SupplierRepository {
	return &supplierRepositoryImpl{
		Database: database,
	}
}

func (repository supplierRepositoryImpl) FindOneByCode(ctx *fasthttp.RequestCtx, supplierCode string) error {
	var supplier string
	query := fmt.Sprintf(`
		select s.code from finance.suppliers s where s.code = '%s'
	`, supplierCode)
	err := repository.Database.QueryRowContext(ctx, query).Scan(&supplier)
	if err == sql.ErrNoRows {
		return errors.New("no data")
	} else {
		return err
	}
}
