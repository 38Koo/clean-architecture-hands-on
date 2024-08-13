package adapterOutPersistance

import "clean-architecture-hands-on/application/domain/model"

type AccountMapper struct {

}

func (am *AccountMapper) MapToDomainEntity(account *AccountEntity, activities []*ActivityEntity, withdrawalBalance, depositBalance int64) model.Account {
	baselineBalance := model.NewMoneyFactory().Subtract(
		model.NewMoney(depositBalance), model.NewMoney(withdrawalBalance))
	
	var accountId model.AccountId = model.AccountId(account.ID)

	return model.WithId(
		&accountId,
		baselineBalance,
		am.MapToActivityWindow(activities),
	)
}

func (am *AccountMapper) MapToActivityWindow(activities []*ActivityEntity) model.ActivityWindow {
	mappedActivities := make([]model.Activity, 0)

	for _, activity := range activities {
		mappedActivities = append(mappedActivities, model.NewActivity(
			model.AccountId(activity.OwnerAccountID),
			model.AccountId(activity.SourceAccountID),
			model.AccountId(activity.TargetAccountID),
			activity.LocalDataTime,
			model.NewMoney(activity.Amount),
		))
	}

	return model.NewActivityWindow(mappedActivities)
}

func (am *AccountMapper) MapToEntity(activity model.Activity) model.Activity {
	return model.NewActivity(
		activity.OwnerAccountId,
		activity.SourceAccountId,
		activity.TargetAccountId,
		activity.Timestamp,
		activity.Money,
	)
}
