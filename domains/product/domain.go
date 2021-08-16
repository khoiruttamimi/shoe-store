package product

import (
	"context"
	"time"
)

type Domain struct {
	ID          int
	Code        string
	Name        string
	Description string
	Price       int
	BrandID     int
	BrandName   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Usecase interface {
	GetAll(ctx context.Context, page, limit int) ([]Domain, int, error)
	GetByID(ctx context.Context, productID int) (Domain, error)
	GetByCode(ctx context.Context, productCode string) (Domain, error)
	Store(ctx context.Context, productDomain *Domain) (Domain, error)
	Update(ctx context.Context, productDomain *Domain) (*Domain, error)
}

type Repository interface {
	GetAll(ctx context.Context, page, limit int) ([]Domain, int, error)
	GetByID(ctx context.Context, productID int) (Domain, error)
	GetByCode(ctx context.Context, productCode string) (Domain, error)
	Store(ctx context.Context, productDomain *Domain) (Domain, error)
	Update(ctx context.Context, productDomain *Domain) (Domain, error)
}
