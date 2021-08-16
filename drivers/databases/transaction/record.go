package transaction

import (
	"shoe-store/domains/transaction"
	"shoe-store/drivers/databases/users"
	"time"
)

type Transaction struct {
	ID         int
	UserID     int
	User       users.Users
	Status     string
	TotalQty   int
	TotalPrice int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func fromDomain(domain *transaction.Domain) *Transaction {
	return &Transaction{
		ID:         domain.ID,
		UserID:     domain.UserID,
		Status:     domain.Status,
		TotalQty:   domain.TotalQty,
		TotalPrice: domain.TotalPrice,
	}
}

func (rec *Transaction) toDomain() transaction.Domain {
	return transaction.Domain{
		ID:         rec.ID,
		UserID:     rec.UserID,
		UserName:   rec.User.Name,
		Status:     rec.Status,
		TotalQty:   rec.TotalQty,
		TotalPrice: rec.TotalPrice,
		CreatedAt:  rec.CreatedAt,
		UpdatedAt:  rec.UpdatedAt,
	}
}
