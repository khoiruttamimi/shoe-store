package request

type TransactionItem struct {
	ProductID int `json:"user_id"`
	Qty       int `json:"qty"`
	Price     int `json:"price"`
}
