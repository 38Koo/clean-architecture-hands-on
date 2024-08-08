package portIn

import (
	"clean-architecture-hands-on/application/domain/model"
	"errors"
)

type SendMoneyCommand struct {
	SourceAccountId model.AccountId
	TargetAccountId model.AccountId
	Money					model.Money
}

func NewSendMoneyCommand(sourceAccountId, targetAccountId model.AccountId, money model.Money) (*SendMoneyCommand, error) {
	if money.IsNegative() {
		return nil, errors.New("the money to send must be positive")
	}

	command := &SendMoneyCommand{
		SourceAccountId: sourceAccountId,
		TargetAccountId: targetAccountId,
		Money: money,
	}

	return command, nil
}
