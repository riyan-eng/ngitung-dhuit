package repository

import (
	"github.com/riyan-eng/ngitung-dhuit/module/finance/repository/model"
	"github.com/riyan-eng/ngitung-dhuit/module/finance/service/entity"
)

type InventoryRepository interface {
	GetByCode(string) (string, error)
	CurrentBalance(string) (model.BalanceInventory, error)
	In(string, entity.InventoryIn) error
	Out() error
}
