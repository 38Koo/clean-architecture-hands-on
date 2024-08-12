package adapterOutPersistance

import (
	"clean-architecture-hands-on/application/domain/model"
	"context"

	"gorm.io/gorm"
)

type IAccountRepository interface {
	Save(ctx context.Context, account *AccountEntity) error
	FindById(ctx context.Context, id uint) (*AccountEntity, error)
}

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (ar *AccountRepository) Save(ctx context.Context, account *AccountEntity) error {
	return ar.db.WithContext(ctx).Save(account).Error
}

func (ar *AccountRepository) FindById(ctx context.Context, id model.AccountId) (*AccountEntity, error) {
	var account AccountEntity
	if err := ar.db.WithContext(ctx).First(&account, id).Error; err != nil {
		return nil, err
	}

	return &account, nil
} 
