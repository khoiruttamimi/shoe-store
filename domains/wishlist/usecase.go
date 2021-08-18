package wishlist

import (
	"context"
	"strings"
	"time"
)

type wishlistUsecase struct {
	wishlistRepository Repository
	contextTimeout     time.Duration
}

func NewWishlistUsecase(timeout time.Duration, cr Repository) Usecase {
	return &wishlistUsecase{
		contextTimeout:     timeout,
		wishlistRepository: cr,
	}
}

func (wu *wishlistUsecase) Store(ctx context.Context, wishlistDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, wu.contextTimeout)
	defer cancel()

	existedWishlist, err := wu.wishlistRepository.CheckWishlist(ctx, wishlistDomain.UserID, wishlistDomain.ProductID)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return Domain{}, err
		}
	}
	if existedWishlist != (Domain{}) {
		wishlistDomain.ID = existedWishlist.ID
		result, err := wu.wishlistRepository.Update(ctx, wishlistDomain)
		if err != nil {
			return Domain{}, err
		}

		return result, nil
	}
	result, err := wu.wishlistRepository.Store(ctx, wishlistDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (wu *wishlistUsecase) GetAll(ctx context.Context, userID int) ([]Domain, error) {
	resp, err := wu.wishlistRepository.GetAll(ctx, userID)
	if err != nil {
		return []Domain{}, err
	}
	return resp, nil
}
