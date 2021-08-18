package response

import (
	"shoe-store/domains/wishlist"
	"time"
)

type Wishlist struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	ProductID   int       `json:"product_id"`
	ProductName string    `json:"product_name"`
	Qty         int       `json:"qty"`
	BasePrice   int       `json:"base_price"`
	TotalPrice  int       `json:"total_price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FromDomain(domain wishlist.Domain) Wishlist {
	return Wishlist{
		ID:          domain.ID,
		UserID:      domain.UserID,
		ProductID:   domain.ProductID,
		ProductName: domain.ProductName,
		Qty:         domain.Qty,
		BasePrice:   domain.BasePrice,
		TotalPrice:  domain.TotalPrice,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
