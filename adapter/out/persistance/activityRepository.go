package adapterOutPersistance

import (
	"clean-architecture-hands-on/application/domain/model"
	"context"
	"time"

	"gorm.io/gorm"
)

type IActivityRepository interface {
	FindByOwnerSince(ctx context.Context, ownerAccountId model.AccountId, since time.Time) ([]*ActivityEntity, error)
	GetDepositBalanceUntil(ctx context.Context, ownerAccountId model.AccountId, until time.Time) (int64, error)
	GetWithdrawalBalanceUntil(ctx context.Context, ownerAccountId model.AccountId, until time.Time) (int64, error)
}

type ActivityRepository struct {
	db *gorm.DB
}

func NewActivityRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (ar *ActivityRepository) FindByOwnerSince(ctx context.Context, ownerAccountId model.AccountId, since time.Time) ([]*ActivityEntity, error) {
	var activities []*ActivityEntity
	if err := ar.db.WithContext(ctx).
		Where("owner_account_id = ? AND timestamp >= ?", ownerAccountId, since).
		Find(&activities).Error; err != nil {
		return nil, err
	}
	return activities, nil
}

func (ar *ActivityRepository) GetDepositBalanceUntil(ctx context.Context, accountId model.AccountId, until time.Time) (int64, error) {
	var sum int64
	if err := ar.db.WithContext(ctx).
		Model(&ActivityEntity{}).
		Select("SUM(amount)").
		Where("target_account_id = ? AND owner_account_id = ? AND timestamp < ?", accountId, accountId, until).
		Scan(&sum).Error; err != nil {
		return 0, err
	}
	return sum, nil
}

func (ar *ActivityRepository) GetWithdrawalBalanceUntil(ctx context.Context, accountId model.AccountId, until time.Time) (int64, error) {
	var sum int64
	if err := ar.db.WithContext(ctx).
		Model(&ActivityEntity{}).
		Select("SUM(amount)").
		Where("source_account_id = ? AND owner_account_id = ? AND timestamp < ?", accountId, accountId, until).
		Scan(&sum).Error; err != nil {
		return 0, err
	}
	return sum, nil
}
