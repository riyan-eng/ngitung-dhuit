package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/riyan-eng/ngitung-dhuit/module/finance/repository/model"
	"github.com/riyan-eng/ngitung-dhuit/module/finance/service/entity"
)

type inventoryRepositoryImpl struct {
	DB *sql.DB
}

func NewInventoryRepository(db *sql.DB) InventoryRepository {
	return &inventoryRepositoryImpl{
		DB: db,
	}
}

func (repository *inventoryRepositoryImpl) GetByCode(code string) (string, error) {
	var goodCode string
	query := fmt.Sprintf(`
		SELECT g.code FROM finance.goods g WHERE g.code = '%s'
	`, code)
	err := repository.DB.QueryRow(query).Scan(&goodCode)
	if err == sql.ErrNoRows {
		return goodCode, errors.New("no data")
	} else {
		return goodCode, err
	}
}

func (repository *inventoryRepositoryImpl) CurrentBalance(good_code string) (model.BalanceInventory, error) {
	var BalanceInventory model.BalanceInventory

	query := fmt.Sprintf(`
		SELECT gs.balance_quantity as quantity, gs.balance_price as price, gs.balance_amount as amount FROM finance.good_stocks gs WHERE gs.good_code='%s' ORDER BY gs.created_at DESC LIMIT 1
	`, good_code)

	err := repository.DB.QueryRow(query).Scan(&BalanceInventory.Quantity, &BalanceInventory.Price, &BalanceInventory.Amount)
	if err == sql.ErrNoRows {
		return BalanceInventory, nil
	}
	return BalanceInventory, err
}

func (repository *inventoryRepositoryImpl) In(transactionID, good_code string, entityInventory entity.InventoryIn) error {
	query := fmt.Sprintf(`
		INSERT INTO finance.good_stocks (transaction_id, good_code, dc, quantity, price, amount, balance_quantity, balance_price, balance_amount)
		VALUES ('%s', '%s', 'D', '%v', '%f', '%f', '%v', '%f', '%f')
	`, transactionID, good_code, entityInventory.Quantity, entityInventory.Price, entityInventory.Amount, entityInventory.BalanceQuantity, entityInventory.BalancePrice, entityInventory.BalanceAmount)
	_, err := repository.DB.Exec(query)
	return err
}

func (repository *inventoryRepositoryImpl) Out() error {
	return nil
}
