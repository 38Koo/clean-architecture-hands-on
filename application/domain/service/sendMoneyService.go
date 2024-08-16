package service

import (
	portIn "clean-architecture-hands-on/application/port/in"
	portOut "clean-architecture-hands-on/application/port/out"
	"fmt"
	"time"
)

type SendMoneyService struct {
	loadAccountPort portOut.LoadAccountPort
	accountLock portOut.AccountLock
	updateAccountStatePort portOut.UpdateAccountState
	moneyTransferProperties MoneyTransferProperties
}

func NewSendMoneyService(
	loadAccountPort portOut.LoadAccountPort,
	accountLock portOut.AccountLock,
	updateAccountStatePort portOut.UpdateAccountState,
	moneyTransferProperties MoneyTransferProperties,
) *SendMoneyService {
	return &SendMoneyService{
		loadAccountPort: loadAccountPort,
		accountLock: accountLock,
		updateAccountStatePort: updateAccountStatePort,
		moneyTransferProperties: moneyTransferProperties,
	}
}

func (s *SendMoneyService) SendMoney(dto portIn.SendMoneyDTO) bool {
	s.checkThreshold(dto)

	baselineDate := time.Now().AddDate(0, 0, -10)

	sourceAccount := s.loadAccountPort.LoadAccount(dto.SourceAccountId, baselineDate)

	targetAccount := s.loadAccountPort.LoadAccount(dto.TargetAccountId, baselineDate)

	sourceAccountId := sourceAccount.GetId()
	if sourceAccountId == nil {
		panic(fmt.Sprintf("expected source account ID not to be empty: %d", dto.SourceAccountId))
	}

	targetAccountId := targetAccount.GetId()
	if targetAccountId == nil {
		panic(fmt.Sprintf("expected target account ID not to be empty: %d", dto.TargetAccountId))
	}

	s.accountLock.AccountLock(*sourceAccountId)
	if (!sourceAccount.Withdraw(dto.Money, *targetAccountId)) {
		s.accountLock.ReleaseAccount(*sourceAccountId)
		return false
	}

	s.accountLock.AccountLock(*targetAccountId)
	if (!targetAccount.Deposit(dto.Money, *sourceAccountId)) {
		s.accountLock.ReleaseAccount(*sourceAccountId)
		s.accountLock.ReleaseAccount(*targetAccountId)
		return false
	}

	s.updateAccountStatePort.UpdateActivities(sourceAccount)
	s.updateAccountStatePort.UpdateActivities(targetAccount)

	s.accountLock.ReleaseAccount(*sourceAccountId)
	s.accountLock.ReleaseAccount(*targetAccountId)
	return true
}

func (s SendMoneyService) checkThreshold(dto portIn.SendMoneyDTO) {
	if dto.Money.IsGreaterThan(s.moneyTransferProperties.maximumTransferThreshold) {
		panic(fmt.Sprintf("the money to send must be less than %s", s.moneyTransferProperties.maximumTransferThreshold))
	}
}