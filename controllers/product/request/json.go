package request

import "shoe-store/domains/product"

type Product struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	BrandID     int    `json:"brand_id"`
}

func (req *Product) ToDomain() *product.Domain {
	return &product.Domain{
		Code:        req.Code,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		BrandID:     req.BrandID,
	}
}
