package repository

type TaxRepository interface {
	GetByCoa(string) (int, error)
}
