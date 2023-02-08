package repository

import "database/sql"

type JournalRepository interface {
	PurchaseJournal()
	SalesJournal()
}

type database struct {
	Journal *sql.DB
}

func NewJournalRepository(db *sql.DB) JournalRepository {
	return &database{
		Journal: db,
	}
}

func (db database) PurchaseJournal() {

}

func (db database) SalesJournal() {

}
