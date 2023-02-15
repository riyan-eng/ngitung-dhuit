package service

import (
	"github.com/riyan-eng/ngitung-dhuit/module/finance/controller/dto"
	"github.com/valyala/fasthttp"
)

type JournalService interface {
	PurchaseJournal(*fasthttp.RequestCtx, *dto.PurchaseJournal) error
	SalesJournal()
	CashPaymentJournal()
}
