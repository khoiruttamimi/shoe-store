package transaction

import (
	"shoe-store/domains/transaction"
	domainTransactionItem "shoe-store/domains/transactionItem"
	"shoe-store/drivers/databases/transactionItem"
	"shoe-store/drivers/databases/users"
	"time"
)

type Transaction struct {
	ID               int
	UserID           int
	User             users.Users
	Status           string
	TotalQty         int
	TotalPrice       int
	TransactionItems []transactionItem.TransactionItem
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func fromDomain(domain *transaction.Domain) *Transaction {
	var items []transactionItem.TransactionItem
	for _, item := range domain.TransactionItems {
		items = append(items, *transactionItem.FromDomain(&item))
	}
	return &Transaction{
		ID:               domain.ID,
		UserID:           domain.UserID,
		Status:           domain.Status,
		TotalQty:         domain.TotalQty,
		TotalPrice:       domain.TotalPrice,
		TransactionItems: items,
	}
}

func (rec *Transaction) toDomain() transaction.Domain {
	var items []domainTransactionItem.Domain
	for _, item := range rec.TransactionItems {
		items = append(items, item.ToDomain())
	}
	return transaction.Domain{
		ID:               rec.ID,
		UserID:           rec.UserID,
		UserName:         rec.User.Name,
		Status:           rec.Status,
		TotalQty:         rec.TotalQty,
		TotalPrice:       rec.TotalPrice,
		TransactionItems: items,
		CreatedAt:        rec.CreatedAt,
		UpdatedAt:        rec.UpdatedAt,
	}
}
