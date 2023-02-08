package repository

type LinkedAccountRepository interface {
	GetByCode(string) (string, error)
}
