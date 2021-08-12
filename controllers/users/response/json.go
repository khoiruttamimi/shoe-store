package response

import (
	"shoe-store/businesses/users"
	"time"
)

type Users struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Username    string    `json:"username"`
	DateOfBirth string    `json:"date_of_birth"`
	Gender      string    `json:"gender"`
	Role        string    `json:"role"`
	Status      bool      `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FromDomain(domain users.Domain) Users {
	return Users{
		Id:          domain.Id,
		Name:        domain.Name,
		Username:    domain.Username,
		DateOfBirth: domain.DateOfBirth,
		Gender:      domain.Gender,
		Role:        domain.Role,
		Status:      domain.Status,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
