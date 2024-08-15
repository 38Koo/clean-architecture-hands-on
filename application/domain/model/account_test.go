package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewAccountWithoutId(t *testing.T) {
	baselineBalance := NewMoney(1000)
	activityWindow := NewActivityWindow([]Activity{})

	account := NewAccountWithoutId(baselineBalance, activityWindow)

	assert.Nil(t, account.GetId(), "account.Id should be nil")
	assert.Equal(t, baselineBalance, account.BaseLineBalance)
	assert.Equal(t, activityWindow, account.ActivityWindow)
}

func TestNewAccountWithId(t *testing.T) {
	accountId := AccountId(1)
	baselineBalance := NewMoney(1000)
	activityWindow := NewActivityWindow([]Activity{})

	account := NewAccountWithId(&accountId, baselineBalance, activityWindow)

	assert.Equal(t, account.GetId(), account.Id)
	assert.Equal(t, baselineBalance, account.BaseLineBalance)
	assert.Equal(t, activityWindow, account.ActivityWindow)
}

func TestCalculateBalance(t *testing.T) {
	accountId := AccountId(1)
	baselineBalance := NewMoney(1000)
	activity1 := NewActivity(accountId, accountId, AccountId(2), time.Now(), NewMoney(200))
	activity2 := NewActivity(accountId, AccountId(3), accountId, time.Now(), NewMoney(300))

	activityWindow := NewActivityWindow([]Activity{activity1, activity2})
	account := NewAccountWithId(&accountId, baselineBalance, activityWindow)

	expectedBalance := NewMoneyFactory().Add(baselineBalance, NewMoney(100))
	calculateBalance := account.CalculateBalance()

	assert.Equal(t, expectedBalance, calculateBalance)
}

func TestWithDraw(t *testing.T) {
	accountId := AccountId(1)
	baseLineBalance := NewMoney(1000)
	activityWindow := NewActivityWindow([]Activity{})
	account := NewAccountWithId(&accountId, baseLineBalance, activityWindow)

	withdrawAmount := NewMoney(500)
	targetAccountId := AccountId(2)

	success := account.Withdraw(withdrawAmount, targetAccountId)

	assert.True(t, success, "Withdrawal should succeed")
	assert.Equal(t, NewMoney(500), account.CalculateBalance(), "Balance should be reduced by withdrawal amount")
}

func TestWithdrawInsufficientBalance(t *testing.T) {
	accountId := AccountId(1)
	baseLineBalance := NewMoney(100)
	activityWindow := NewActivityWindow([]Activity{})
	account := NewAccountWithId(&accountId, baseLineBalance, activityWindow)

	withdrawAmount := NewMoney(200)
	targetAccountId := AccountId(2)

	success := account.Withdraw(withdrawAmount, targetAccountId)

	assert.False(t, success, "Withdrawal should fail due to insufficient balance")
	assert.Equal(t, baseLineBalance, account.CalculateBalance(), "Balance should remain unchanged")
}

func TestDeposit(t *testing.T) {
	accountId := AccountId(1)
	baseLineBalance := NewMoney(1000)
	activityWindow := NewActivityWindow([]Activity{})
	account := NewAccountWithId(&accountId, baseLineBalance, activityWindow)

	depositAmount := NewMoney(500)
	sourceAccountId := AccountId(2)

	success := account.Deposit(depositAmount, sourceAccountId)

	assert.True(t, success, "Deposit should succeed")
	assert.Equal(t, NewMoney(1500), account.CalculateBalance(), "Balance should be increased by deposit amount")
}
