package response

import (
	"shoe-store/domains/users"
)

type AuthResponse struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Role  string `json:"role"`
	Token string `json:"token"`
}

func GetAuthResponse(domain users.Domain, token string) AuthResponse {
	return AuthResponse{
		Name:  domain.Name,
		Phone: domain.Phone,
		Role:  domain.Role,
		Token: token,
	}
}
