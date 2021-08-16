package response

import (
	"shoe-store/domains/transactionItem"
	"time"
)

type TransactionItem struct {
	ID            int       `json:"id"`
	TransactionID int       `json:"transaction_id"`
	ProductID     int       `json:"product_id"`
	ProductName   string    `json:"product_name"`
	Qty           int       `json:"qty"`
	BasePrice     int       `json:"base_price"`
	TotalPrice    int       `json:"total_price"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func FromDomain(domain transactionItem.Domain) TransactionItem {
	return TransactionItem{
		ID:            domain.ID,
		TransactionID: domain.TransactionID,
		ProductID:     domain.ProductID,
		ProductName:   domain.ProductName,
		Qty:           domain.Qty,
		BasePrice:     domain.BasePrice,
		TotalPrice:    domain.TotalPrice,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdatedAt,
	}
}
