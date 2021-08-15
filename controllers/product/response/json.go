package response

import (
	"shoe-store/domains/product"
	"time"
)

type Product struct {
	ID          int       `json:"id"`
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	BrandID     int       `json:"brand_id"`
	BrandName   string    `json:"brand_name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FromDomain(domain product.Domain) Product {
	return Product{
		ID:          domain.ID,
		Code:        domain.Code,
		Name:        domain.Name,
		Description: domain.Description,
		Price:       domain.Price,
		BrandID:     domain.BrandID,
		BrandName:   domain.BrandName,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
