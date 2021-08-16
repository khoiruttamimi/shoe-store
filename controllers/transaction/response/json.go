package response

import (
	"shoe-store/controllers/transactionItem/response"
	"shoe-store/domains/transaction"
	"time"
)

type Transaction struct {
	ID               int                        `json:"id"`
	UserID           int                        `json:"user_id"`
	UserName         string                     `json:"user_name"`
	Status           string                     `json:"status"`
	TotalQty         int                        `json:"total_qty"`
	TotalPrice       int                        `json:"total_price"`
	TransactionItems []response.TransactionItem `json:"transaction_items"`
	CreatedAt        time.Time                  `json:"created_at"`
	UpdatedAt        time.Time                  `json:"updated_at"`
}

func FromDomain(domain transaction.Domain) Transaction {
	var items []response.TransactionItem
	for _, item := range domain.TransactionItems {
		items = append(items, response.FromDomain(item))
	}
	return Transaction{
		ID:               domain.ID,
		UserID:           domain.UserID,
		UserName:         domain.UserName,
		Status:           domain.Status,
		TotalQty:         domain.TotalQty,
		TotalPrice:       domain.TotalPrice,
		TransactionItems: items,
		CreatedAt:        domain.CreatedAt,
		UpdatedAt:        domain.UpdatedAt,
	}
}
