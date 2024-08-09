package service

import portOut "clean-architecture-hands-on/application/port/out"

type SendMoneyService struct {
	loadAccountPort portOut.LoadAccountPort
	accountLock portOut.AccountLock
}

func NewSendMoneyService(
loadAccountPort portOut.LoadAccountPort,
accountLock portOut.AccountLock,
) *SendMoneyService {
	return &SendMoneyService{
		loadAccountPort: loadAccountPort,
		accountLock: accountLock,
	}
}
