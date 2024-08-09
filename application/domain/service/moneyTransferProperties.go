package service

import "clean-architecture-hands-on/application/domain/model"

type MoneyTransferProperties struct {
	maximumTransferThreshold model.Money
}

func NewMoneyTransferProperties() MoneyTransferProperties {
	return MoneyTransferProperties{
		maximumTransferThreshold: model.NewMoney(1_000_000),
	}
}

func (m *MoneyTransferProperties) GetMaximumTransferThreshold() model.Money {
	return m.maximumTransferThreshold
}

func (m *MoneyTransferProperties) SetMaximumTransferThreshold(maximumTransferThreshold model.Money) {
	m.maximumTransferThreshold = maximumTransferThreshold
}
