package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/valyala/fasthttp"
)

type supplierRepositoryImpl struct {
	DB *sql.DB
}

func NewSupplierRepository(db *sql.DB) SupplierRepository {
	return &supplierRepositoryImpl{
		DB: db,
	}
}

func (repository supplierRepositoryImpl) FindOne(ctx *fasthttp.RequestCtx, supplierCode string) error {
	var supplier string
	query := fmt.Sprintf(`
		select s.code from finance.suppliers s where s.code = '%s'
	`, supplierCode)
	err := repository.DB.QueryRowContext(ctx, query).Scan(&supplier)
	if err == sql.ErrNoRows {
		return errors.New("no data")
	} else {
		return err
	}
}
