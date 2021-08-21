package wishlist

import (
	"shoe-store/domains/wishlist"
	"shoe-store/drivers/databases/product"
	"shoe-store/drivers/databases/users"

	"time"

	"gorm.io/gorm"
)

type Wishlist struct {
	ID        int
	UserID    int
	User      users.Users
	ProductID int
	Product   product.Product
	Qty       int
	CreatedAt time.Time
	UpdatedAt time.Time
	gorm.DeletedAt
}

func FromDomain(domain *wishlist.Domain) *Wishlist {
	return &Wishlist{
		ID:        domain.ID,
		UserID:    domain.UserID,
		ProductID: domain.ProductID,
		Qty:       domain.Qty,
	}
}
func (rec *Wishlist) ToDomain() wishlist.Domain {
	return wishlist.Domain{
		ID:          rec.ID,
		UserID:      rec.UserID,
		ProductID:   rec.ProductID,
		ProductName: rec.Product.Name,
		Qty:         rec.Qty,
		BasePrice:   rec.Product.Price,
		TotalPrice:  rec.Product.Price * rec.Qty,
		CreatedAt:   rec.CreatedAt,
		UpdatedAt:   rec.UpdatedAt,
	}
}
