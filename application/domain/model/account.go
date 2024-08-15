package model

import (
	"time"
)

type AccountId int32

type Account struct {
	Id              *AccountId
	BaseLineBalance Money
	ActivityWindow  ActivityWindow
}

func NewAccountWithoutId(baseLineBalance Money, activityWindow ActivityWindow) Account {
	return Account{
		Id:              nil,
		BaseLineBalance: baseLineBalance,
		ActivityWindow:  activityWindow,
	}
}

func NewAccountWithId(accountId *AccountId, baseLineBalance Money, activityWindow ActivityWindow) Account {
	return Account{
		Id:              accountId,
		BaseLineBalance: baseLineBalance,
		ActivityWindow:  activityWindow,
	}
}

func (a Account) GetId() *AccountId {
	return a.Id
}

func (a Account) CalculateBalance() Money {
	moneyFactory := NewMoneyFactory()
	return moneyFactory.Add(a.BaseLineBalance, a.ActivityWindow.CalculateBalance(*a.Id))
}

func (a *Account) Withdraw(money Money, targetAccountId AccountId) bool {
	if !a.mayWithdrawal(money) {
		return false
	}

	withdrawal := NewActivity(*a.Id, *a.Id, targetAccountId, time.Now(), money)
	a.ActivityWindow.AddActivity(withdrawal)
	return true
}

func (a Account) mayWithdrawal(money Money) bool {
	moneyFactory := NewMoneyFactory()
	return moneyFactory.Add(a.CalculateBalance(), money.Negate()).IsPositiveOrZero()
}

func (a *Account) Deposit(money Money, sourceAccountId AccountId) bool {
	deposit := NewActivity(*a.Id, sourceAccountId, *a.Id, time.Now(), money)
	a.ActivityWindow.AddActivity(deposit)
	return true
}
