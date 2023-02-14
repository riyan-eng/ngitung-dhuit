package repository

import (
	"github.com/riyan-eng/ngitung-dhuit/module/finance/service/entity"
	"github.com/valyala/fasthttp"
)

type JournalRepository interface {
	PurchaseJournal(*fasthttp.RequestCtx, entity.PurchaseJournal) error
	SalesJournal()
}
