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

type Transaction struct {
	Coa    string  `json:"coa" validate:"required"`
	Amount float64 `json:"amount" validate:"required"`
}

type CashPaymentJournal struct {
	Debet  []Transaction `json:"debet" validate:"required"`
	Credit []Transaction `json:"credit" validate:"required"`
}
