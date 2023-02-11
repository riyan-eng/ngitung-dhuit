package repository

type SubsidiaryLedgerRepository interface {
	InsertPayable(string, float64) error
}
