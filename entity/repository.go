package entity

type TransactionRepository interface {
	Insert(id string, AccountID string, amount float64, status string, errorMessage string) error
}
