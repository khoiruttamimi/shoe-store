package users

import (
	"shoe-store/businesses/users"
	"time"
)

type Users struct {
	ID          int
	Name        string
	Username    string
	Password    string
	DateOfBirth string
	Gender      string
	Role        string
	Status      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (rec *Users) toDomain() users.Domain {
	return users.Domain{
		Id:          rec.ID,
		Name:        rec.Name,
		Username:    rec.Username,
		Password:    rec.Password,
		DateOfBirth: rec.DateOfBirth,
		Gender:      rec.Gender,
		Role:        rec.Role,
		Status:      rec.Status,
		CreatedAt:   rec.CreatedAt,
		UpdatedAt:   rec.UpdatedAt,
	}
}

func fromDomain(userDomain users.Domain) *Users {
	return &Users{
		ID:          userDomain.Id,
		Name:        userDomain.Name,
		Username:    userDomain.Username,
		Password:    userDomain.Password,
		DateOfBirth: userDomain.DateOfBirth,
		Gender:      userDomain.Gender,
		Role:        userDomain.Role,
		Status:      userDomain.Status,
		CreatedAt:   userDomain.CreatedAt,
		UpdatedAt:   userDomain.UpdatedAt,
	}
}
