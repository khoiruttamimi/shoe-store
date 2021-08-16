package users

import (
	"shoe-store/domains/users"
	"time"
)

type Users struct {
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

func (rec *Users) toDomain() users.Domain {
	return users.Domain{
		ID:          rec.ID,
		Name:        rec.Name,
		DateOfBirth: rec.DateOfBirth,
		Gender:      rec.Gender,
		Address:     rec.Address,
		Phone:       rec.Phone,
		Password:    rec.Password,
		Role:        rec.Role,
		Status:      rec.Status,
		CreatedAt:   rec.CreatedAt,
		UpdatedAt:   rec.UpdatedAt,
	}
}

func fromDomain(userDomain users.Domain) *Users {
	return &Users{
		ID:          userDomain.ID,
		Name:        userDomain.Name,
		DateOfBirth: userDomain.DateOfBirth,
		Gender:      userDomain.Gender,
		Address:     userDomain.Address,
		Phone:       userDomain.Phone,
		Password:    userDomain.Password,
		Role:        userDomain.Role,
		Status:      userDomain.Status,
		CreatedAt:   userDomain.CreatedAt,
		UpdatedAt:   userDomain.UpdatedAt,
	}
}
