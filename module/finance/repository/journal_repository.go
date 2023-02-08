package repository

import "github.com/riyan-eng/ngitung-dhuit/module/finance/service/entity"

type JournalRepository interface {
	PurchaseJournal(entity.PurchaseJournal) error
	SalesJournal()
}
