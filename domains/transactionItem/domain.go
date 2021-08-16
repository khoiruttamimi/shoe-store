package transactionItem

import (
	"time"
)

type Domain struct {
	ID            int
	TransactionID int
	ProductID     int
	ProductName   string
	Qty           int
	BasePrice     int
	TotalPrice    int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
