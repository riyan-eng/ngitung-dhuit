package repository

type TaxRepository interface {
	GetByCoa(string) (string, error)
}
