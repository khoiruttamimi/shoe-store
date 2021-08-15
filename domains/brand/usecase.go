package brand

import (
	"context"
	"shoe-store/domains"
	"strings"
	"time"
)

type brandUsecase struct {
	brandRepository Repository
	contextTimeout  time.Duration
}

func NewBrandUsecase(timeout time.Duration, cr Repository) Usecase {
	return &brandUsecase{
		contextTimeout:  timeout,
		brandRepository: cr,
	}
}

func (bu *brandUsecase) Store(ctx context.Context, brandDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()

	existedBrand, err := bu.brandRepository.GetByName(ctx, brandDomain.Name)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return Domain{}, err
		}
	}
	if existedBrand != (Domain{}) {
		return Domain{}, domains.ErrDuplicateData
	}

	result, err := bu.brandRepository.Store(ctx, brandDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (bu *brandUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	resp, err := bu.brandRepository.Find(ctx, "")
	if err != nil {
		return []Domain{}, err
	}
	return resp, nil
}

func (bu *brandUsecase) GetByID(ctx context.Context, id int) (Domain, error) {
	if id <= 0 {
		return Domain{}, domains.ErrIDNotFound
	}

	resp, err := bu.brandRepository.FindByID(id)
	if err != nil {
		return Domain{}, err
	}
	return resp, nil
}
