package repository

import "database/sql"

type inventoryRepositoryImpl struct {
	DB *sql.DB
}

func NewInventoryRepository(db *sql.DB) InventoryRepository {
	return &inventoryRepositoryImpl{
		DB: db,
	}
}

func (repository *inventoryRepositoryImpl) GetByCode(code string) (string, error) {
	return "", nil
}
