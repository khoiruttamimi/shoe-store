package brand

import (
	"context"
	"shoe-store/domains/brand"

	"gorm.io/gorm"
)

type brandRepository struct {
	conn *gorm.DB
}

func NewBrandRepository(conn *gorm.DB) brand.Repository {
	return &brandRepository{
		conn: conn,
	}
}

func (br *brandRepository) Find(ctx context.Context, active string) ([]brand.Domain, error) {
	rec := []Brand{}

	err := br.conn.Find(&rec).Error
	if err != nil {
		return []brand.Domain{}, err
	}

	brandDomain := []brand.Domain{}
	for _, value := range rec {
		brandDomain = append(brandDomain, value.toDomain())
	}

	return brandDomain, nil
}

func (br *brandRepository) GetByName(ctx context.Context, brandName string) (brand.Domain, error) {
	rec := Brand{}
	err := br.conn.Where("name = ?", brandName).First(&rec).Error
	if err != nil {
		return brand.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (br *brandRepository) Store(ctx context.Context, brandDomain *brand.Domain) (brand.Domain, error) {
	rec := fromDomain(brandDomain)

	result := br.conn.Create(&rec)
	if result.Error != nil {
		return brand.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (br *brandRepository) FindByID(id int) (brand.Domain, error) {
	rec := Brand{}

	if err := br.conn.Where("id = ?", id).First(&rec).Error; err != nil {
		return brand.Domain{}, err
	}
	return rec.toDomain(), nil
}
