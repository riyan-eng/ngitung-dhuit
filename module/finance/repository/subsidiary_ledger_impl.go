package repository

import (
	"database/sql"
	"fmt"
)

type subsidiaryLedgerRepositoryImpl struct {
	DB *sql.DB
}

func NewSubsidiaryLedgerRepository(db *sql.DB) SubsidiaryLedgerRepository {
	return &subsidiaryLedgerRepositoryImpl{
		DB: db,
	}
}

func (repository subsidiaryLedgerRepositoryImpl) InsertPayable(supplier_code, transactionID string, amount float64) error {
	query := fmt.Sprintf(`
		INSERT INTO finance.account_payable_subsidiary_ledgers (supplier_code, transaction_id, dc, amount)
		VALUES ('%s', '%s', 'C', '%f')
	`, supplier_code, transactionID, amount)

	_, err := repository.DB.Exec(query)
	return err
}
