package adapterOutPersistance

import (
	"clean-architecture-hands-on/application/domain/model"
	"context"
	"errors"
	"time"
)


type AccountPersistenceAdapter struct {
	accountRepository *AccountRepository
	activityRepository *ActivityRepository
	accountMapper *AccountMapper
}

func NewAccountPersistenceAdapter(
	accountRepository *AccountRepository,
	activityRepository *ActivityRepository,
	accountMapper *AccountMapper,
) *AccountPersistenceAdapter {
	return &AccountPersistenceAdapter{
		accountRepository: accountRepository,
		activityRepository: activityRepository,
		accountMapper: accountMapper,
	}
}

func (apa *AccountPersistenceAdapter) LoadAccount(accountId model.AccountId, baselineDate time.Time) (*model.Account, error) {
	ctx := context.Background()
	accountEntity, err := apa.accountRepository.FindById(ctx, accountId)
	if err != nil {
		return nil, errors.New("entity not found")
	}

	activitiesEntity, err := apa.activityRepository.FindByOwnerSince(ctx, accountId, baselineDate)
	if err != nil {
		return nil, err
	}

	withdrawalBalance, err := apa.activityRepository.GetWithdrawalBalanceUntil(ctx, accountId, baselineDate)
	if err != nil {
		return nil, err
	}

	depositBalance, err := apa.activityRepository.GetDepositBalanceUntil(ctx, accountId, baselineDate)
	if err != nil {
		return nil, err
	}

	account := apa.accountMapper.MapToDomainEntity(accountEntity, activitiesEntity, withdrawalBalance, depositBalance)
	return &account, nil
}

func (apa *AccountPersistenceAdapter) UpdateActivities(account model.Account) {
	for _, activity := range account.ActivityWindow.GetActivities() {
		if activity.Id == nil {
			apa.activityRepository.db.Save(apa.accountMapper.MapToEntity(activity))
		}
	}
}
