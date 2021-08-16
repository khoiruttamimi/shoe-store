package users

import (
	"context"
	"shoe-store/app/middleware"
	"shoe-store/domains"
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

func (uc *userUsecase) Register(ctx context.Context, userDomain *Domain) (Domain, string, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	existedUser, err := uc.userRepository.GetByPhone(ctx, userDomain.Phone)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return Domain{}, "", err
		}
	}
	if existedUser != (Domain{}) {
		return Domain{}, "", domains.ErrDuplicateData
	}

	userDomain.Password, err = encrypt.Hash(userDomain.Password)
	userDomain.Role = "customer"
	userDomain.Status = true
	if err != nil {
		return Domain{}, "", domains.ErrInternalServer
	}
	user, err := uc.userRepository.Store(ctx, userDomain)
	if err != nil {
		return Domain{}, "", err
	}
	token := uc.jwtAuth.GenerateToken(user.ID)
	return user, token, nil
}

func (uc *userUsecase) Login(ctx context.Context, userDomain *Domain) (Domain, string, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if strings.TrimSpace(userDomain.Phone) == "" || strings.TrimSpace(userDomain.Password) == "" {
		return Domain{}, "", domains.ErrPhonePasswordNotFound
	}

	user, err := uc.userRepository.GetByPhone(ctx, userDomain.Phone)
	if err != nil {
		return Domain{}, "", err
	}

	if !encrypt.ValidateHash(userDomain.Password, user.Password) {
		return Domain{}, "", domains.ErrInvalidCredential
	}

	token := uc.jwtAuth.GenerateToken(user.ID)
	return user, token, nil
}

func (cu *userUsecase) GetProfile(ctx context.Context, userID int) (Domain, error) {
	if userID <= 0 {
		return Domain{}, domains.ErrIDNotFound
	}

	user, err := cu.userRepository.GetByID(ctx, userID)
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (cu *userUsecase) UpdateProfile(ctx context.Context, userDomain *Domain) (Domain, error) {
	_, err := cu.userRepository.GetByID(ctx, userDomain.ID)
	if err != nil {
		return Domain{}, err
	}
	result, err := cu.userRepository.Update(ctx, userDomain)
	if err != nil {
		return Domain{}, err
	}
	return result, nil
}
