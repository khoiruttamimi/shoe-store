package transaction

import (
	"context"
	"shoe-store/domains/transaction"

	"gorm.io/gorm"
)

type transactionRepository struct {
	conn *gorm.DB
}

func NewTransactionRepository(conn *gorm.DB) transaction.Repository {
	return &transactionRepository{
		conn: conn,
	}
}

func (tr *transactionRepository) GetAll(ctx context.Context, userID int) ([]transaction.Domain, error) {
	rec := []Transaction{}

	err := tr.conn.Preload("User").Where("user_id = ?", userID).Find(&rec).Error
	if err != nil {
		return []transaction.Domain{}, err
	}

	transactionDomain := []transaction.Domain{}
	for _, value := range rec {
		transactionDomain = append(transactionDomain, value.toDomain())
	}

	return transactionDomain, nil
}

func (tr *transactionRepository) GetByID(ctx context.Context, transactionID int, userID int) (transaction.Domain, error) {
	rec := Transaction{}
	err := tr.conn.
		Preload("User").
		Preload("TransactionItems").
		Preload("TransactionItems.Product").
		Where("id = ?", transactionID).
		Where("user_id = ?", userID).
		First(&rec).Error
	if err != nil {
		return transaction.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (tr *transactionRepository) Store(ctx context.Context, transactionDomain *transaction.Domain) (transaction.Domain, error) {
	rec := fromDomain(transactionDomain)

	result := tr.conn.Create(&rec)
	if result.Error != nil {
		return transaction.Domain{}, result.Error
	}
	err := tr.conn.Preload("User").Preload("TransactionItems").Preload("TransactionItems.Product").First(&rec, rec.ID).Error
	if err != nil {
		return transaction.Domain{}, result.Error
	}
	return rec.toDomain(), nil
}
