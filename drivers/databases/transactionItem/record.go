package transactionItem

import (
	"shoe-store/domains/transactionItem"
	"shoe-store/drivers/databases/product"

	"time"
)

type TransactionItem struct {
	ID            int
	TransactionID int
	ProductID     int
	Product       product.Product
	Qty           int
	Price         int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func FromDomain(domain *transactionItem.Domain) *TransactionItem {
	return &TransactionItem{
		TransactionID: domain.TransactionID,
		ProductID:     domain.ProductID,
		Qty:           domain.Qty,
		Price:         domain.BasePrice,
	}
}
func (rec *TransactionItem) ToDomain() transactionItem.Domain {
	return transactionItem.Domain{
		ID:            rec.ID,
		TransactionID: rec.TransactionID,
		ProductID:     rec.ProductID,
		ProductName:   rec.Product.Name,
		Qty:           rec.Qty,
		BasePrice:     rec.Price,
		TotalPrice:    rec.Price * rec.Qty,
		CreatedAt:     rec.CreatedAt,
		UpdatedAt:     rec.UpdatedAt,
	}
}
