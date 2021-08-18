package wishlist

import (
	"context"
	"time"
)

type Domain struct {
	ID          int
	UserID      int
	ProductID   int
	ProductName string
	Qty         int
	BasePrice   int
	TotalPrice  int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
type Usecase interface {
	Store(ctx context.Context, wishlistDomain *Domain) (Domain, error)
	GetAll(ctx context.Context, userID int) ([]Domain, error)
}

type Repository interface {
	GetAll(ctx context.Context, userID int) ([]Domain, error)
	CheckWishlist(ctx context.Context, UserID, ProductID int) (Domain, error)
	Store(ctx context.Context, wishlistDomain *Domain) (Domain, error)
	Update(ctx context.Context, wishlistDomain *Domain) (Domain, error)
}
