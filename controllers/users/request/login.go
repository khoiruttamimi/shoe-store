package request

import "shoe-store/domains/users"

type UserLogin struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func (req *UserLogin) ToDomain() *users.Domain {
	return &users.Domain{
		Phone:    req.Phone,
		Password: req.Password,
	}
}
