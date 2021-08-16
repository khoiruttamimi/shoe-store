package request

import "shoe-store/domains/transactionItem"

type TransactionItem struct {
	ProductID    int `json:"product_id"`
	Qty          int `json:"qty"`
	ProductPrice int `json:"product_price"`
}

func (req *TransactionItem) ToDomain() *transactionItem.Domain {
	return &transactionItem.Domain{
		ProductID: req.ProductID,
		Qty:       req.Qty,
		BasePrice: req.ProductPrice,
	}
}
