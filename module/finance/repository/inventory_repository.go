package repository

type InventoryRepository interface {
	GetByCode(code string) (string, error)
}
