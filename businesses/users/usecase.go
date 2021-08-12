package users

import (
	"context"
	"shoe-store/app/middleware"
	"shoe-store/businesses"
	"shoe-store/helpers/encrypt"
	"strings"
	"time"
)

type userUsecase struct {
	userRepository Repository
	contextTimeout time.Duration
	jwtAuth        *middleware.ConfigJWT
}

func NewUserUsecase(ur Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration) Usecase {
	return &userUsecase{
		userRepository: ur,
		jwtAuth:        jwtauth,
		contextTimeout: timeout,
	}
}

func (uc *userUsecase) CreateToken(ctx context.Context, username, password string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if strings.TrimSpace(username) == "" && strings.TrimSpace(password) == "" {
		return "", businesses.ErrUsernamePasswordNotFound
	}

	userDomain, err := uc.userRepository.GetByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	if !encrypt.ValidateHash(password, userDomain.Password) {
		return "", businesses.ErrInternalServer
	}

	token := uc.jwtAuth.GenerateToken(userDomain.Id)
	return token, nil
}

func (uc *userUsecase) Store(ctx context.Context, userDomain *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	existedUser, err := uc.userRepository.GetByUsername(ctx, userDomain.Username)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return err
		}
	}
	if existedUser != (Domain{}) {
		return businesses.ErrDuplicateData
	}

	userDomain.Password, err = encrypt.Hash(userDomain.Password)
	if err != nil {
		return businesses.ErrInternalServer
	}
	err = uc.userRepository.Store(ctx, userDomain)
	if err != nil {
		return err
	}

	return nil
}
func (uc *userUsecase) Login(ctx context.Context, userDomain *Domain) (interface{}, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if strings.TrimSpace(userDomain.Username) == "" && strings.TrimSpace(userDomain.Password) == "" {
		return nil, businesses.ErrUsernamePasswordNotFound
	}

	user, err := uc.userRepository.GetByUsername(ctx, userDomain.Username)
	if err != nil {
		return nil, err
	}

	if !encrypt.ValidateHash(userDomain.Password, user.Password) {
		return nil, businesses.ErrInternalServer
	}

	token := uc.jwtAuth.GenerateToken(user.Id)
	res := struct {
		Role  string `json:"role"`
		Token string `json:"token"`
	}{Role: user.Role, Token: token}

	return res, nil
}

func (cu *userUsecase) GetProfile(ctx context.Context, userID int) (Domain, error) {
	if userID <= 0 {
		return Domain{}, businesses.ErrIDNotFound
	}

	resp, err := cu.userRepository.GetByID(ctx, userID)
	if err != nil {
		return Domain{}, err
	}
	return resp, nil
}
