package repository

import "database/sql"

type taxRepositoryImpl struct {
	DB *sql.DB
}

func NewTaxRepository(db *sql.DB) TaxRepository {
	return &taxRepositoryImpl{
		DB: db,
	}
}

func (repository *taxRepositoryImpl) GetByCoa(string) (int, error) {
	return 11, nil
}
