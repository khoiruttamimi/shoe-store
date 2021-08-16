package brand

import (
	"context"
	"time"
)

type Domain struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	GetByID(ctx context.Context, id int) (Domain, error)
	Store(ctx context.Context, brandDomain *Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
}

type Repository interface {
	Find(ctx context.Context, active string) ([]Domain, error)
	FindByID(id int) (Domain, error)
	GetByName(ctx context.Context, brandName string) (Domain, error)
	Store(ctx context.Context, brandDomain *Domain) (Domain, error)
}
