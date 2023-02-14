package repository

import (
	"database/sql"
	"errors"
	"fmt"
)

type supplierRepositoryImpl struct {
	DB *sql.DB
}

func NewSupplierRepository(db *sql.DB) SupplierRepository {
	return &supplierRepositoryImpl{
		DB: db,
	}
}

func (repository supplierRepositoryImpl) FindOne(supplierCode string) error {
	var supplier string
	query := fmt.Sprintf(`
		select s.code from finance.suppliers s where s.code = '%s'
	`, supplierCode)
	err := repository.DB.QueryRow(query).Scan(&supplier)
	if err == sql.ErrNoRows {
		return errors.New("no data")
	} else {
		return err
	}
}
