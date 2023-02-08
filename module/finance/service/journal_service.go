package service

import (
	"github.com/riyan-eng/ngitung-dhuit/module/finance/controller/dto"
)

type JournalService interface {
	PurchaseJournal(*dto.PurchaseJournal) error
	SalesJournal()
}
