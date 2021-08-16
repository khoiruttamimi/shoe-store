package request

import (
	"shoe-store/controllers/transactionItem/request"
	"shoe-store/domains/transaction"
)

type Transaction struct {
	UserID int                       `json:"user_id"`
	Status string                    `json:"status"`
	Items  []request.TransactionItem `json:"items"`
}

func (req *Transaction) ToDomain() *transaction.Domain {
	var totalQty int
	var totalPrice int

	for _, item := range req.Items {
		totalQty += item.Qty
		totalPrice += item.Price
	}

	return &transaction.Domain{
		UserID:     req.UserID,
		Status:     req.Status,
		TotalQty:   totalQty,
		TotalPrice: totalPrice,
	}
}
