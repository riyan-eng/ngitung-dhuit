package entity

type PurchaseJournal struct {
	TransactionID string
	Debet         PurchaseJournalDebet
	Credit        PurchaseJournalCredit
}

type PurchaseJournalDebet struct {
	Code   string
	Amount float64
}

type PurchaseJournalCredit struct {
	Code   string
	Amount float64
}
