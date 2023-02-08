package repository

import "database/sql"

type linkedAccountRepositoryImpl struct {
	DB *sql.DB
}

func NewLinkedAccountRepository(db *sql.DB) LinkedAccountRepository {
	return &linkedAccountRepositoryImpl{
		DB: db,
	}
}

func (repository *linkedAccountRepositoryImpl) GetByCode(string) (string, error) {
	return "", nil
}
