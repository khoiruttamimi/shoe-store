package request

import (
	"shoe-store/controllers/transactionItem/request"
	"shoe-store/domains/transaction"
	"shoe-store/domains/transactionItem"
)

type Transaction struct {
	Items []request.TransactionItem `json:"items"`
}

func (req *Transaction) ToDomain() *transaction.Domain {
	var totalQty int
	var totalPrice int

	var items []transactionItem.Domain

	for _, item := range req.Items {
		totalQty += item.Qty
		totalPrice += item.ProductPrice * item.Qty
		items = append(items, *item.ToDomain())
	}

	return &transaction.Domain{
		TotalQty:         totalQty,
		TotalPrice:       totalPrice,
		TransactionItems: items,
	}
}
