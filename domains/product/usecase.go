package product

import (
	"context"
	"shoe-store/domains"
	"shoe-store/domains/brand"
	"shoe-store/domains/iplocator"
	"strings"
	"time"
)

type productUsecase struct {
	productRepository Repository
	brandUsecase      brand.Usecase
	contextTimeout    time.Duration
	ipLocator         iplocator.Repository
}

func NewProductUsecase(nr Repository, cu brand.Usecase, timeout time.Duration, il iplocator.Repository) Usecase {
	return &productUsecase{
		productRepository: nr,
		brandUsecase:      cu,
		contextTimeout:    timeout,
		ipLocator:         il,
	}
}

func (pu *productUsecase) GetAll(ctx context.Context, page, limit int) ([]Domain, int, error) {
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	defer cancel()

	res, total, err := pu.productRepository.GetAll(ctx, page, limit)
	if err != nil {
		return []Domain{}, 0, err
	}

	return res, total, nil
}

func (pu *productUsecase) GetByID(ctx context.Context, productID int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	defer cancel()

	if productID <= 0 {
		return Domain{}, domains.ErrProductIDResource
	}
	res, err := pu.productRepository.GetByID(ctx, productID)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (pu *productUsecase) GetByCode(ctx context.Context, productCode string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	defer cancel()

	if strings.TrimSpace(productCode) == "" {
		return Domain{}, domains.ErrProductCodeResource
	}
	res, err := pu.productRepository.GetByCode(ctx, productCode)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}
func (pu *productUsecase) Store(ctx context.Context, productDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	defer cancel()

	_, err := pu.brandUsecase.GetByID(ctx, productDomain.BrandID)
	if err != nil {
		return Domain{}, domains.ErrBrandNotFound
	}

	existedProduct, err := pu.productRepository.GetByCode(ctx, productDomain.Code)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return Domain{}, err
		}
	}
	if existedProduct != (Domain{}) {
		return Domain{}, domains.ErrDuplicateData
	}

	result, err := pu.productRepository.Store(ctx, productDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}
func (pu *productUsecase) Update(ctx context.Context, productDomain *Domain) (*Domain, error) {
	_, err := pu.productRepository.GetByID(ctx, productDomain.ID)
	if err != nil {
		return &Domain{}, domains.ErrIDNotFound
	}
	existedProduct, err := pu.productRepository.GetByCode(ctx, productDomain.Code)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return &Domain{}, err
		}
	}
	if existedProduct != (Domain{}) && existedProduct.ID != productDomain.ID {
		return &Domain{}, domains.ErrDuplicateData
	}
	result, err := pu.productRepository.Update(ctx, productDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}
