package repository

import "database/sql"

func NewJournalRepository(db *sql.DB) JournalRepository {
	return &journalRepositoryImpl{
		Database: db,
	}
}

type journalRepositoryImpl struct {
	Database *sql.DB
}

func (repository *journalRepositoryImpl) PurchaseJournal() {

}

func (repository *journalRepositoryImpl) SalesJournal() {

}
