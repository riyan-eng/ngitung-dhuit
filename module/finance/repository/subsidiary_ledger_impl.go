package repository

import (
	"database/sql"
	"fmt"

	"github.com/valyala/fasthttp"
)

type subsidiaryLedgerRepositoryImpl struct {
	DB *sql.DB
}

func NewSubsidiaryLedgerRepository(db *sql.DB) SubsidiaryLedgerRepository {
	return &subsidiaryLedgerRepositoryImpl{
		DB: db,
	}
}

func (repository subsidiaryLedgerRepositoryImpl) InsertPayable(ctx *fasthttp.RequestCtx, supplier_code, transactionID string, amount float64) error {
	query := fmt.Sprintf(`
		INSERT INTO finance.account_payable_subsidiary_ledgers (supplier_code, transaction_id, dc, amount)
		VALUES ('%s', '%s', 'C', '%f')
	`, supplier_code, transactionID, amount)

	_, err := repository.DB.ExecContext(ctx, query)
	return err
}
