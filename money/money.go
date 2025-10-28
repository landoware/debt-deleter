package money

import (
	"fmt"
)

type Money struct {
	Cents int
}

func NewMoney(dollars, cents int) *Money {
	return &Money{
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
// Comparison
func (m *Money) LessThan(comparend *Money) bool {
	return m.Cents < comparend.Cents
}

func (m *Money) LT(comparend *Money) bool {
	return m.LessThan(comparend)

}

func (m *Money) LessThanOrEqualTo(comparend *Money) bool {
	return m.Cents <= comparend.Cents
}

func (m *Money) LTE(comparend *Money) bool {
	return m.LessThanOrEqualTo(comparend)
}

func (m *Money) GreaterThan(comparend *Money) bool {
	return m.Cents > comparend.Cents
}

func (m *Money) GT(comparend *Money) bool {
	return m.GreaterThan(comparend)
}

func (m *Money) GreaterThanOrEqualTo(comparend *Money) bool {
	return m.Cents >= comparend.Cents
}

func (m *Money) GTE(comparend *Money) bool {
	return m.GreaterThanOrEqualTo(comparend)
}

func (m *Money) GreaterThanZero() bool {
	return m.Cents > 0
}

func (m *Money) GreaterThanOrEqualToZero() bool {
	return m.Cents >= 0
}

func (m *Money) LessThanZero() bool {
	return m.Cents < 0
}

func (m *Money) LessThanOrEqualToZero() bool {
	return m.Cents <= 0
}
