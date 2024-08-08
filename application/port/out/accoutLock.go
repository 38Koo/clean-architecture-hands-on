package portOut

import "clean-architecture-hands-on/application/domain/model"

type AccountLock interface {
	AccountLock(accountId model.AccountId) error
}