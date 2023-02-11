package repository

import (
	"database/sql"
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
	query := fmt.Sprintf(`
		SELECT * FROM finance.goods g WHERE g.code = '%s'
	`, code)
	err := repository.DB.QueryRow(query).Err()
	return "", err
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

func (repository *inventoryRepositoryImpl) In(good_code string, entityInventory entity.InventoryIn) error {
	query := fmt.Sprintf(`
		INSERT INTO finance.good_stocks (good_code, dc, quantity, price, amount, balance_quantity, balance_price, balance_amount)
		VALUES ('%s', 'D', '%v', '%f', '%f', '%v', '%f', '%f')
	`, good_code, entityInventory.Quantity, entityInventory.Price, entityInventory.Amount, entityInventory.BalanceQuantity, entityInventory.BalancePrice, entityInventory.BalanceAmount)
	_, err := repository.DB.Exec(query)
	return err
}

func (repository *inventoryRepositoryImpl) Out() error {
	return nil
}
