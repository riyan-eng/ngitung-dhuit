package repository

type TransactionRepository interface {
	Insert(string, float64) (string, error)
}
