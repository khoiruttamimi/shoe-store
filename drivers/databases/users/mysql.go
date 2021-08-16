package users

import (
	"context"
	"shoe-store/domains/users"

	"gorm.io/gorm"
)

type mysqlUsersRepository struct {
	Conn *gorm.DB
}

func NewMySQLUserRepository(conn *gorm.DB) users.Repository {
	return &mysqlUsersRepository{
		Conn: conn,
	}
}

func (ur *mysqlUsersRepository) Fetch(ctx context.Context, page, perpage int) ([]users.Domain, int, error) {
	rec := []Users{}

	offset := (page - 1) * perpage
	err := ur.Conn.Offset(offset).Limit(perpage).Find(&rec).Error
	if err != nil {
		return []users.Domain{}, 0, err
	}

	var totalData int64
	err = ur.Conn.Count(&totalData).Error
	if err != nil {
		return []users.Domain{}, 0, err
	}

	var domainProduct []users.Domain
	for _, value := range rec {
		domainProduct = append(domainProduct, value.toDomain())
	}
	return domainProduct, int(totalData), nil
}

func (ur *mysqlUsersRepository) GetByID(ctx context.Context, userID int) (users.Domain, error) {
	rec := Users{}
	err := ur.Conn.Where("id = ?", userID).First(&rec).Error
	if err != nil {
		return users.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (ur *mysqlUsersRepository) GetByPhone(ctx context.Context, phone string) (users.Domain, error) {
	rec := Users{}
	err := ur.Conn.Where("phone = ?", phone).First(&rec).Error
	if err != nil {
		return users.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (ur *mysqlUsersRepository) Store(ctx context.Context, userDomain *users.Domain) (users.Domain, error) {
	rec := fromDomain(*userDomain)

	result := ur.Conn.Create(&rec)
	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (ur *mysqlUsersRepository) Update(ctx context.Context, userDomain *users.Domain) (users.Domain, error) {
	rec := fromDomain(*userDomain)

	result := ur.Conn.Updates(&rec)
	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}
