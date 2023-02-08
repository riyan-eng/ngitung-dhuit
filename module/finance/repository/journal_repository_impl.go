package repository

import (
	"database/sql"

	"github.com/riyan-eng/ngitung-dhuit/module/finance/service/entity"
)

func NewJournalRepository(db *sql.DB) JournalRepository {
	return &journalRepositoryImpl{
		Database: db,
	}
}

type journalRepositoryImpl struct {
	Database *sql.DB
}

func (repository *journalRepositoryImpl) PurchaseJournal(journal entity.PurchaseJournal) error {
	return nil
}

func (repository *journalRepositoryImpl) SalesJournal() {

}
