package debts

import (
	"github.com/landoware/debt-deleter/interest"
	"github.com/landoware/debt-deleter/money"
)

type Loan struct {
	Name           string
	Principal      money.Money
	UnpaidInterest money.Money
	Rate           interest.Rate
	MinPayment     money.Money
	DueDay         int
}

func NewLoan(name string, principal money.Money, rate interest.Rate, min_payment money.Money, due_day int, unpaidInterest money.Money) Loan {
	return Loan{
		Name:           name,
		Principal:      principal,
		UnpaidInterest: unpaidInterest,
		Rate:           rate,
		MinPayment:     min_payment,
		DueDay:         due_day,
	}
}

func (loan *Loan) PayOnLoan(amount money.Money) {
	// Reduce interest
	if loan.UnpaidInterest.GreaterThanOrEqualTo(amount) {
		amount = amount.Subtract(loan.UnpaidInterest)
		loan.UnpaidInterest.Cents = 0
	}

	// Allocate the rest to principal
	if amount.GreaterThanZero() {
		loan.Principal = loan.Principal.Subtract(amount)
	}
}
