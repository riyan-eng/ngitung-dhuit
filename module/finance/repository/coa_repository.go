package repository

type COARepository interface {
	GetByCode(string) error
}
