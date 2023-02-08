package service

import "github.com/riyan-eng/ngitung-dhuit/module/finance/repository"

type JournalService interface {
	PurchaseJournal()
	SalesJournal()
}

type journalRepository struct {
	Journal repository.JournalRepository
}

func NewJournalSerice(journal repository.JournalRepository) JournalService {
	return &journalRepository{
		Journal: journal,
	}
}

func (repo journalRepository) PurchaseJournal() {

}

func (repo journalRepository) SalesJournal() {

}
