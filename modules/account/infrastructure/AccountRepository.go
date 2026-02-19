package infrastructure

import (
	"context"
	"errors"

	"gorm.io/gorm"

	helper "UnpakSiamida/common/helper"
	domain "UnpakSiamida/modules/account/domain"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) domain.IAccountRepository {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) Auth(ctx context.Context, username string, password string) (*domain.Account, error) {
	var user domain.Account

	// 1️⃣ Cari berdasarkan username saja
	err := r.db.WithContext(ctx).
		Where("username = ?", username).
		Where("aktif = 'Y'").
		First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	hashed := helper.Sha1MD5(password)

	if user.Password != hashed {
		return nil, gorm.ErrRecordNotFound
	}

	return &user, nil
}

func (r *AccountRepository) Get(ctx context.Context, userUUID string) (*domain.Account, error) {
	// var user domain.Account

	// err := r.db.WithContext(ctx).
	// 	Where("uuid = ?", userUUID).
	// 	First(&user).Error
	// if err != nil {
	// 	return nil, err
	// }

	return nil, errors.New("featur not implemented")
}

func (r *AccountRepository) GetByUserId(ctx context.Context, userid string) (*domain.Account, error) {
	var user domain.Account

	err := r.db.WithContext(ctx).
		Where("userid = ?", userid).
		Where("aktif = 'Y'").
		First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
