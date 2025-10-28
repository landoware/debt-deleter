package money

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
