package repository

import (
	"github.com/riyan-eng/ngitung-dhuit/module/finance/repository/model"
	"github.com/riyan-eng/ngitung-dhuit/module/finance/service/entity"
	"github.com/valyala/fasthttp"
)

type InventoryRepository interface {
	GetByCode(*fasthttp.RequestCtx, string) (string, error)
	CurrentBalance(*fasthttp.RequestCtx, string) (model.BalanceInventory, error)
	In(*fasthttp.RequestCtx, string, string, entity.InventoryIn) error
	Out() error
}
