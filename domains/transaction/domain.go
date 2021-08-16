package transaction

import (
	"context"
	"shoe-store/domains/transactionItem"
	"time"
)

type Domain struct {
	ID               int
	UserID           int
	UserName         string
	Status           string
	TotalQty         int
	TotalPrice       int
	TransactionItems []transactionItem.Domain
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Usecase interface {
	Store(ctx context.Context, transactionDomain *Domain) (Domain, error)
	GetAll(ctx context.Context, userID int) ([]Domain, error)
	GetByID(ctx context.Context, transactionID int, userID int) (Domain, error)
}

type Repository interface {
	Store(ctx context.Context, transactionDomain *Domain) (Domain, error)
	GetAll(ctx context.Context, userID int) ([]Domain, error)
	GetByID(ctx context.Context, transactionID int, userID int) (Domain, error)
}
