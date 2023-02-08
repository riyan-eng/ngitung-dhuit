package entity

type PurchaseJournal struct {
	Debet  PurchaseJournalDebet
	Credit PurchaseJournalCredit
}

type PurchaseJournalDebet struct {
	Code   string
	Amount float64
}

type PurchaseJournalCredit struct {
	Code   string
	Amount float64
}
