package entity

type InventoryIn struct {
	Quantity        int
	Price           float64
	Amount          float64
	BalanceQuantity int
	BalancePrice    float64
	BalanceAmount   float64
}

type Inventory struct {
	Quantity int
	Price    float64
	Amount   float64
}
