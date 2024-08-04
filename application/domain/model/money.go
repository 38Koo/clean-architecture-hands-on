package model

import "math/big"

type Money struct {
	Amount *big.Int
}

var ZERO = NewMoney(0)

func NewMoney(value int64) Money {
	return Money{Amount: big.NewInt(value)}
}

func (m Money) IsPositiveOrZero() bool {
	return m.Amount.Cmp(big.NewInt(0)) >= 0
}

func (m Money) IsNegative() bool {
	return m.Amount.Cmp(big.NewInt(0)) < 0
}

func (m Money) IsPositive() bool {
	return m.Amount.Cmp(big.NewInt(0)) > 0
}

func (m Money) IsGreaterThanOrEqualTo(money Money) bool {
	return m.Amount.Cmp(money.Amount) >= 0
}

func (m Money) IsGreaterThan(money Money) bool {
	return m.Amount.Cmp(money.Amount) > 0
}

func Add(originMoney, targetMoney Money) Money {
	return Money{Amount: new(big.Int).Add(originMoney.Amount, targetMoney.Amount)}
}

func (m Money) Minus(money Money) Money {
	return Money{Amount: new(big.Int).Sub(m.Amount, money.Amount)}
}

func (m Money) Plus(money Money) Money {
	return Money{Amount: new(big.Int).Add(m.Amount, money.Amount)}
}

func Subtract(originMoney, targetMoney Money) Money {
	return Money{Amount: new(big.Int).Sub(originMoney.Amount, targetMoney.Amount)}
}

func (m Money) Negate() Money {
	return Money{Amount: new(big.Int).Neg(m.Amount)}
}
