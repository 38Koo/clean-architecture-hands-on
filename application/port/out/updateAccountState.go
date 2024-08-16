package portOut

import "clean-architecture-hands-on/application/domain/model"

type UpdateAccountState interface {
	UpdateActivities(account model.Account)
}