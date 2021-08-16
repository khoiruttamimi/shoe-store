package transaction

import (
	"context"
	"shoe-store/domains"
	"time"
)

type transactionUsecase struct {
	transactionRepository Repository
	contextTimeout        time.Duration
}

func NewTransactionUsecase(timeout time.Duration, cr Repository) Usecase {
	return &transactionUsecase{
		contextTimeout:        timeout,
		transactionRepository: cr,
	}
}

func (tu *transactionUsecase) Store(ctx context.Context, transactionDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, tu.contextTimeout)
	defer cancel()

	transactionDomain.Status = "paid"
	result, err := tu.transactionRepository.Store(ctx, transactionDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (tu *transactionUsecase) GetAll(ctx context.Context, userID int) ([]Domain, error) {
	resp, err := tu.transactionRepository.GetAll(ctx, userID)
	if err != nil {
		return []Domain{}, err
	}
	return resp, nil
}

func (tu *transactionUsecase) GetByID(ctx context.Context, transactionID int, userID int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, tu.contextTimeout)
	defer cancel()

	if transactionID <= 0 {
		return Domain{}, domains.ErrProductIDResource
	}
	res, err := tu.transactionRepository.GetByID(ctx, transactionID, userID)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}
