package request

import "shoe-store/domains/users"

type UpdateProfile struct {
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	DateOfBirth string `json:"date_of_birth"`
	Gender      string `json:"gender"`
	Address     string `json:"address"`
}

func (req *UpdateProfile) ToDomain() *users.Domain {
	return &users.Domain{
		Name:        req.Name,
		Phone:       req.Phone,
		DateOfBirth: req.DateOfBirth,
		Gender:      req.Gender,
		Address:     req.Address,
	}
}
