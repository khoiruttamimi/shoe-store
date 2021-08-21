package wishlist

import (
	"context"
	"shoe-store/domains/wishlist"

	"gorm.io/gorm"
)

type wishlistRepository struct {
	conn *gorm.DB
}

func NewWishlistRepository(conn *gorm.DB) wishlist.Repository {
	return &wishlistRepository{
		conn: conn,
	}
}

func (wr *wishlistRepository) GetAll(ctx context.Context, userID int) ([]wishlist.Domain, error) {
	rec := []Wishlist{}

	err := wr.conn.Preload("Product").Where("user_id = ?", userID).Find(&rec).Error
	if err != nil {
		return []wishlist.Domain{}, err
	}

	wishlistDomain := []wishlist.Domain{}
	for _, value := range rec {
		wishlistDomain = append(wishlistDomain, value.ToDomain())
	}

	return wishlistDomain, nil
}

func (wr *wishlistRepository) CheckWishlist(ctx context.Context, UserID, ProductID int) (wishlist.Domain, error) {
	rec := Wishlist{}
	err := wr.conn.Where("user_id = ?", UserID).Where("product_id = ?", ProductID).First(&rec).Error
	if err != nil {
		return wishlist.Domain{}, err
	}
	return rec.ToDomain(), nil
}

func (wr *wishlistRepository) Store(ctx context.Context, wishlistDomain *wishlist.Domain) (wishlist.Domain, error) {
	rec := FromDomain(wishlistDomain)

	result := wr.conn.Create(&rec)
	if result.Error != nil {
		return wishlist.Domain{}, result.Error
	}
	err := wr.conn.Preload("Product").First(&rec, rec.ID).Error
	if err != nil {
		return wishlist.Domain{}, result.Error
	}
	return rec.ToDomain(), nil
}

func (wr *wishlistRepository) Update(ctx context.Context, wishlistDomain *wishlist.Domain) (wishlist.Domain, error) {
	rec := FromDomain(wishlistDomain)

	result := wr.conn.Updates(&rec)
	if result.Error != nil {
		return wishlist.Domain{}, result.Error
	}
	err := wr.conn.Preload("Product").First(&rec, rec.ID).Error
	if err != nil {
		return wishlist.Domain{}, result.Error
	}
	return rec.ToDomain(), nil
}

func (wr *wishlistRepository) Delete(ctx context.Context, wishlistDomain *wishlist.Domain) (wishlist.Domain, error) {
	rec := FromDomain(wishlistDomain)

	result := wr.conn.Delete(&rec)
	if result.Error != nil {
		return wishlist.Domain{}, result.Error
	}
	return rec.ToDomain(), nil
}
