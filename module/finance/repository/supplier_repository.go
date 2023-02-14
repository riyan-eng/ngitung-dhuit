package repository

type SupplierRepository interface {
	FindOne(string) error
}
