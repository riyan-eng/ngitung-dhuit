package repository

type SubsidiaryLedgerRepository interface {
	InsertPayable(string, string, float64) error
}
