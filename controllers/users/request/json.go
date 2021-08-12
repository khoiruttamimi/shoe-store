package request

import "shoe-store/businesses/users"

type Users struct {
	Name        string `json:"name"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	DateOfBirth string `json:"date_of_birth"`
	Gender      string `json:"gender"`
	Role        string `json:"role"`
	Status      bool   `json:"status"`
}

func (req *Users) ToDomain() *users.Domain {
	return &users.Domain{
		Name:        req.Name,
		Username:    req.Username,
		Password:    req.Password,
		DateOfBirth: req.DateOfBirth,
		Gender:      req.Gender,
		Role:        req.Role,
		Status:      req.Status,
	}
}
