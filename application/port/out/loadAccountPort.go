package portOut

import (
	"clean-architecture-hands-on/application/domain/model"
	"time"
)

type LoadAccountPort interface {
	LoadAccount(accountId model.AccountId, localDateTime time.Time) model.Account
}
