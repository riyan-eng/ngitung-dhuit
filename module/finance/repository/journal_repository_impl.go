package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/riyan-eng/ngitung-dhuit/module/finance/service/entity"
	"github.com/valyala/fasthttp"
)

func NewJournalRepository(db *sql.DB) JournalRepository {
	return &journalRepositoryImpl{
		Database: db,
	}
}

type journalRepositoryImpl struct {
	Database *sql.DB
}

func (repository *journalRepositoryImpl) PurchaseJournal(ctx *fasthttp.RequestCtx, journal entity.PurchaseJournal) error {

	queryPurchaseJournalDebet := fmt.Sprintf(`
		INSERT INTO finance.purchase_journals (transaction_id, coa_code, dc, amount) VALUES ('%s', '%s', 'D', '%f')
	`, journal.TransactionID, journal.Debet.Code, journal.Debet.Amount)

	queryPurchaseJournalCredit := fmt.Sprintf(`
		INSERT INTO finance.purchase_journals (transaction_id, coa_code, dc, amount) VALUES ('%s', '%s', 'C', '%f')
	`, journal.TransactionID, journal.Credit.Code, journal.Credit.Amount)

	tx, err := repository.Database.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, queryPurchaseJournalDebet)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, queryPurchaseJournalCredit)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return err
}

func (repository *journalRepositoryImpl) SalesJournal() {

}
