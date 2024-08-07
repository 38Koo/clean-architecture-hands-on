package model

import "math/big"

type MoneyFactory struct{}

func NewMoneyFactory() *MoneyFactory {
	return &MoneyFactory{}
}

func (mf *MoneyFactory) Add(originMoney, targetMoney Money) Money {
	return Money{Amount: new(big.Int).Add(originMoney.Amount, targetMoney.Amount)}
}

func (mf *MoneyFactory) Subtract(originMoney, targetMoney Money) Money {
	return Money{Amount: new(big.Int).Sub(originMoney.Amount, targetMoney.Amount)}
}
