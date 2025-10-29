package money

import (
	"fmt"
)

type Money struct {
	Cents int
}

func NewMoney(dollars, cents int) Money {
	return Money{
		Cents: dollars*100 + cents,
	}
}

func (m Money) Dollars() int {
	return m.Cents / 100
}

func (m Money) OnlyCents() int {
	return m.Cents % 100
}

func (m Money) String() string {
	return fmt.Sprintf("$%d.%d", m.Dollars(), m.OnlyCents())
}

// Arithmetic
func (m Money) Add(other Money) Money {
	return Money{m.Cents + other.Cents}
}

func (m Money) Subtract(other Money) Money {
	return Money{m.Cents - other.Cents}
}

// Comparison
func (m Money) Equals(comparend Money) bool {
	return m.Cents == comparend.Cents
}

func (m Money) NotEquals(comparend Money) bool {
	return m.Cents != comparend.Cents
}

func (m Money) LessThan(comparend Money) bool {
	return m.Cents < comparend.Cents
}

func (m Money) LessThanOrEqualTo(comparend Money) bool {
	return m.Cents <= comparend.Cents
}

func (m Money) LTE(comparend Money) bool {
	return m.LessThanOrEqualTo(comparend)
}

func (m Money) GreaterThan(comparend Money) bool {
	return m.Cents > comparend.Cents
}

func (m Money) GreaterThanOrEqualTo(comparend Money) bool {
	return m.Cents >= comparend.Cents
}

func (m Money) GreaterThanZero() bool {
	return m.Cents > 0
}

func (m Money) GreaterThanOrEqualToZero() bool {
	return m.Cents >= 0
}

func (m Money) LessThanZero() bool {
	return m.Cents < 0
}

func (m Money) LessThanOrEqualToZero() bool {
	return m.Cents <= 0
}
