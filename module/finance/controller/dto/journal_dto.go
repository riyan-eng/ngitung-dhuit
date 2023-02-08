package dto

type PurchaseJournal struct {
	InventoryCode string  `json:"inventory_code" validate:"required"`
	Quantity      uint    `json:"quantity" validate:"required"`
	Price         float64 `json:"price" validate:"required"`
	PPNIncome     bool    `json:"ppn_income" validate:"required"`
	FreightPaid   float64 `json:"freight_paid"`
}

type SalesJournal struct {
	InventoryCode    string  `json:"inventory_code" validate:"required"`
	Quantity         uint    `json:"quantity" validate:"required"`
	Price            float64 `json:"price" validate:"required"`
	PPNOutcome       bool    `json:"ppn_outcome" validate:"required"`
	FreightCollected float64 `json:"freight_collected"`
}
