package request

import "shoe-store/domains/users"

type UserRegister struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func (req *UserRegister) ToDomain() *users.Domain {
	return &users.Domain{
		Name:     req.Name,
		Phone:    req.Phone,
		Password: req.Password,
	}
}
