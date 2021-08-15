package transaction

import (
	"context"
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

func (bu *transactionUsecase) Store(ctx context.Context, transactionDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()

	result, err := bu.transactionRepository.Store(ctx, transactionDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (bu *transactionUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	resp, err := bu.transactionRepository.GetAll(ctx, "paid")
	if err != nil {
		return []Domain{}, err
	}
	return resp, nil
}
