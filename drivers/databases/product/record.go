package product

import (
	productUsecase "shoe-store/domains/product"
	"shoe-store/drivers/databases/brand"
	"time"
)

type Product struct {
	ID          int
	Code        string
	Name        string
	Description string
	Price       int
	BrandID     int
	Brand       brand.Brand
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func fromDomain(domain *productUsecase.Domain) *Product {
	return &Product{
		ID:          domain.ID,
		Code:        domain.Code,
		Name:        domain.Name,
		Description: domain.Description,
		Price:       domain.Price,
		BrandID:     domain.BrandID,
	}
}

func (rec *Product) toDomain() productUsecase.Domain {
	return productUsecase.Domain{
		ID:          rec.ID,
		Code:        rec.Code,
		Name:        rec.Name,
		Description: rec.Description,
		Price:       rec.Price,
		BrandID:     rec.BrandID,
		BrandName:   rec.Brand.Name,
		CreatedAt:   rec.CreatedAt,
		UpdatedAt:   rec.UpdatedAt,
	}
}
