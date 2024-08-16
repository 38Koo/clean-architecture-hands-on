package service

import (
	"clean-architecture-hands-on/application/domain/model"
	portIn "clean-architecture-hands-on/application/port/in"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockLoadAccountPort struct {
	mock.Mock
}

func (m *MockLoadAccountPort) LoadAccount(accountId model.AccountId, baselineDate time.Time) model.Account {
	args := m.Called(accountId, baselineDate)
	return args.Get(0).(model.Account)
}

type MockAccountLock struct {
	mock.Mock
}

func (m *MockAccountLock) AccountLock(accountId model.AccountId) error {
	m.Called(accountId)
	return nil
}

func (m *MockAccountLock) ReleaseAccount(accountId model.AccountId) error {
	m.Called(accountId)
	return nil
}

type MockUpdateAccountState struct {
	mock.Mock
}

func (m *MockUpdateAccountState) UpdateActivities(account model.Account) {
	m.Called(account)
}

func TestTransactionSucceeds(t *testing.T) {
	mockLoadAccountPort := new(MockLoadAccountPort)
	mockAccountLock := new(MockAccountLock)
	mockUpdateAccountState := new(MockUpdateAccountState)
	moneyTransferProperties := MoneyTransferProperties{
		maximumTransferThreshold: model.NewMoney(1000),
	}

	sendMoneyService := NewSendMoneyService(
		mockLoadAccountPort,
		mockAccountLock,
		mockUpdateAccountState,
		moneyTransferProperties,
	)

	sourceAccountId := model.AccountId(1)
	targetAccountId := model.AccountId(2)
	baselineBalance := model.NewMoney(1000)
	activityWindow := model.NewActivityWindow([]model.Activity{})

	sourceAccount := model.NewAccountWithId(&sourceAccountId, baselineBalance, activityWindow)
	targetAccount := model.NewAccountWithId(&targetAccountId, baselineBalance, activityWindow)

	sourceAccount.Withdraw(model.NewMoney(500), targetAccountId)
	targetAccount.Deposit(model.NewMoney(500), sourceAccountId)

	money := model.NewMoney(500)

	// Mock LoadAccount responses
	mockLoadAccountPort.On("LoadAccount", sourceAccountId, mock.AnythingOfType("time.Time")).Return(sourceAccount)
	mockLoadAccountPort.On("LoadAccount", targetAccountId, mock.AnythingOfType("time.Time")).Return(targetAccount)

	// Mock AccountLock
	mockAccountLock.On("AccountLock", sourceAccountId).Return()
	mockAccountLock.On("AccountLock", targetAccountId).Return()
	mockAccountLock.On("ReleaseAccount", sourceAccountId).Return()
	mockAccountLock.On("ReleaseAccount", targetAccountId).Return()

	// Mock UpdateActivities with more flexible matching
	mockUpdateAccountState.On("UpdateActivities", mock.MatchedBy(func(account model.Account) bool {
		return len(account.ActivityWindow.GetActivities()) == 2
	})).Return()

	dto := portIn.SendMoneyDTO{
		SourceAccountId: *sourceAccount.GetId(),
		TargetAccountId: *targetAccount.GetId(),
		Money:          money,
	}

	isSuccess := sendMoneyService.SendMoney(dto)

	assert.True(t, isSuccess)

	mockAccountLock.AssertCalled(t, "AccountLock", sourceAccountId)
	mockAccountLock.AssertCalled(t, "ReleaseAccount", sourceAccountId)

	mockAccountLock.AssertCalled(t, "AccountLock", targetAccountId)
	mockAccountLock.AssertCalled(t, "ReleaseAccount", targetAccountId)

	thenAccountsHaveBeenUpdated(t, mockUpdateAccountState, sourceAccountId, targetAccountId)
}

func thenAccountsHaveBeenUpdated(t *testing.T, mockUpdateAccountState *MockUpdateAccountState, accountIds ...model.AccountId) {
	accountCaptor := mockUpdateAccountState.Calls

	// アカウントIDのリストを取得
	var updatedAccountIds []model.AccountId
	for _, call := range accountCaptor {
			if call.Method == "UpdateActivities" {
					account := call.Arguments.Get(0).(model.Account)
					updatedAccountIds = append(updatedAccountIds, *account.GetId())
			}
	}

	// 各アカウントIDが更新されたことを確認
	for _, accountId := range accountIds {
			assert.Contains(t, updatedAccountIds, accountId)
	}
}
