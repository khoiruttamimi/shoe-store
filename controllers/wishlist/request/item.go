package request

import "shoe-store/domains/wishlist"

type Wishlist struct {
	ProductID int `json:"product_id"`
	Qty       int `json:"qty"`
}

func (req *Wishlist) ToDomain() *wishlist.Domain {
	return &wishlist.Domain{
		ProductID: req.ProductID,
		Qty:       req.Qty,
	}
}
