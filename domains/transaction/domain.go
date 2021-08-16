package transaction

import (
	"context"
	"time"
)

type Domain struct {
	ID         int
	UserID     int
	UserName   string
	Status     string
	TotalQty   int
	TotalPrice int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Usecase interface {
	Store(ctx context.Context, transactionDomain *Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
}

type Repository interface {
	Store(ctx context.Context, transactionDomain *Domain) (Domain, error)
	GetAll(ctx context.Context, status string) ([]Domain, error)
}
