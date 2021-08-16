package brand

import (
	"shoe-store/domains/brand"
	"time"
)

type Brand struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func fromDomain(domain *brand.Domain) *Brand {
	return &Brand{
		ID:   domain.ID,
		Name: domain.Name,
	}
}

func (rec *Brand) toDomain() brand.Domain {
	return brand.Domain{
		ID:        rec.ID,
		Name:      rec.Name,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}
