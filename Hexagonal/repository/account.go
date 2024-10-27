package repository

type Account struct {
	AccountID   int     `gorm:"primaryKey;column:account_id;autoIncrement"`
	CustomerID  int     `gorm:"column:customer_id"`
	OpeningDate string  `gorm:"column:opening_date"`
	AccountType string  `gorm:"column:account_type"`
	Amount      float64 `gorm:"column:amount"`
	Status      int     `gorm:"column:status"`
}

// port
type AccountRpeository interface {
	Create(Account) (*Account , error)
	GetAll(int) ([]Account , error)
}