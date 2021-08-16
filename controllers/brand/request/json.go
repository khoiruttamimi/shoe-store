package request

import "shoe-store/domains/brand"

type Brand struct {
	Name string `json:"name"`
}

func (req *Brand) ToDomain() *brand.Domain {
	return &brand.Domain{
		Name: req.Name,
	}
}
