package repository

import (
	"database/sql"
	"fmt"
)

func NewCOARepository(db *sql.DB) COARepository {
	return &chartOfAccountRepositoryImpl{
		Database: db,
	}
}

type chartOfAccountRepositoryImpl struct {
	Database *sql.DB
}

func (repository *chartOfAccountRepositoryImpl) GetByCode(coa string) error {
	query := fmt.Sprintf(`
		select coa.id from master.coa coa where coa.code = '%v'
	`, coa)

	err := repository.Database.QueryRow(query).Err()
	if err != nil {
		return err
	}
	return nil
}
