package portIn

import (
	"clean-architecture-hands-on/application/domain/model"
	"errors"
)

type SendMoneyDTO struct {
	SourceAccountId model.AccountId
	TargetAccountId model.AccountId
	Money						model.Money
}

func NewSendMoneyCommand(sourceAccountId, targetAccountId model.AccountId, money model.Money) (*SendMoneyDTO, error) {
	if err := validateInput(sourceAccountId, targetAccountId, money); err != nil {
		return nil, err
	}

	sendMoneyDTO := &SendMoneyDTO{
		SourceAccountId: sourceAccountId,
		TargetAccountId: targetAccountId,
		Money: money,
	}

	return sendMoneyDTO, nil
}

func validateInput(sourceAccountId, targetAccountId model.AccountId, money model.Money) error {
	if money.IsNegative() {
		return errors.New("the money to send must be positive")
	}

	return nil
}
