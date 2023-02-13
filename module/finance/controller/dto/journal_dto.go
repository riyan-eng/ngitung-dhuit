package dto

type Goods struct {
	GoodCode string  `json:"good_code" validate:"required"`
	Qty      int     `json:"quantity" validate:"required"`
	Price    float64 `json:"price" validate:"required"`
	Discount int     `json:"discount"`
}

type PurchaseJournal struct {
	Goods         []Goods `json:"goods"`
	PPNIncome     bool    `json:"ppn_income"`
	FreightPaid   float64 `json:"freight_paid"`
	CreditAccount string  `json:"credit_account" validate:"required"`
	SupplierCode  string  `json:"supplier_code"`
	Description   string  `json:"description"`
}

type SalesJournal struct {
	InventoryCode    string  `json:"inventory_code" validate:"required"`
	Quantity         uint    `json:"quantity" validate:"required"`
	Price            float64 `json:"price" validate:"required"`
	PPNOutcome       bool    `json:"ppn_outcome"`
	FreightCollected float64 `json:"freight_collected"`
	DebitAccount     string  `json:"debit_account" validate:"required"`
}
