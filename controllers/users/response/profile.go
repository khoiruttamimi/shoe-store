package response

import (
	"shoe-store/domains/users"
	"time"
)

type ProfileResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Phone       string    `json:"phone"`
	DateOfBirth string    `json:"date_of_birth"`
	Gender      string    `json:"gender"`
	Address     string    `json:"address"`
	Role        string    `json:"role"`
	Status      bool      `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func GetProfileResponse(domain users.Domain) ProfileResponse {
	return ProfileResponse{
		ID:          domain.ID,
		Name:        domain.Name,
		Phone:       domain.Phone,
		DateOfBirth: domain.DateOfBirth,
		Gender:      domain.Gender,
		Address:     domain.Address,
		Role:        domain.Role,
		Status:      domain.Status,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
