package users

import (
	"context"
	"time"
)

type Domain struct {
	ID          int
	Name        string
	DateOfBirth string
	Gender      string
	Address     string
	Phone       string
	Password    string
	Role        string
	Status      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Usecase interface {
	Register(ctx context.Context, data *Domain) (Domain, string, error)
	Login(ctx context.Context, data *Domain) (Domain, string, error)
	GetProfile(ctx context.Context, userID int) (Domain, error)
	UpdateProfile(ctx context.Context, req *Domain) (Domain, error)
}

type Repository interface {
	GetByID(ctx context.Context, userID int) (Domain, error)
	GetByPhone(ctx context.Context, phone string) (Domain, error)
	Store(ctx context.Context, data *Domain) (Domain, error)
	Update(ctx context.Context, data *Domain) (Domain, error)
}
