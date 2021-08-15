package product

import (
	"context"
	"shoe-store/domains/product"

	"gorm.io/gorm"
)

type mysqlProductRepository struct {
	Conn *gorm.DB
}

func NewMySQLProductRepository(conn *gorm.DB) product.Repository {
	return &mysqlProductRepository{
		Conn: conn,
	}
}

func (pr *mysqlProductRepository) GetAll(ctx context.Context, page, limit int) ([]product.Domain, int, error) {
	rec := []Product{}

	offset := (page - 1) * limit
	err := pr.Conn.Preload("Brand").Offset(offset).Limit(limit).Find(&rec).Error
	if err != nil {
		return []product.Domain{}, 0, err
	}

	var totalData int64
	err = pr.Conn.Model(&rec).Count(&totalData).Error
	if err != nil {
		return []product.Domain{}, 0, err
	}

	var domainProduct []product.Domain
	for _, value := range rec {
		domainProduct = append(domainProduct, value.toDomain())
	}
	return domainProduct, int(totalData), nil
}

func (pr *mysqlProductRepository) GetByID(ctx context.Context, productID int) (product.Domain, error) {
	rec := Product{}
	err := pr.Conn.Preload("Brand").Where("id = ?", productID).First(&rec).Error
	if err != nil {
		return product.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (pr *mysqlProductRepository) GetByCode(ctx context.Context, productCode string) (product.Domain, error) {
	rec := Product{}
	err := pr.Conn.Where("code = ?", productCode).First(&rec).Error
	if err != nil {
		return product.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (pr *mysqlProductRepository) Store(ctx context.Context, productDomain *product.Domain) (product.Domain, error) {
	rec := fromDomain(productDomain)

	result := pr.Conn.Create(&rec)
	if result.Error != nil {
		return product.Domain{}, result.Error
	}

	err := pr.Conn.Preload("Brand").First(&rec, rec.ID).Error
	if err != nil {
		return product.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (pr *mysqlProductRepository) Update(ctx context.Context, productDomain *product.Domain) (product.Domain, error) {
	rec := fromDomain(productDomain)

	result := pr.Conn.Updates(&rec)
	if result.Error != nil {
		return product.Domain{}, result.Error
	}

	err := pr.Conn.Preload("Brand").First(&rec, rec.ID).Error
	if err != nil {
		return product.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}
