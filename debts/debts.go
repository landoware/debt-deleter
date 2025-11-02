package debts

import (
	"github.com/landoware/debt-deleter/interest"
	"github.com/landoware/debt-deleter/money"
	"github.com/uniplaces/carbon"
)

type Loan struct {
	Name           string
	Principal      money.Money
	UnpaidInterest money.Money
	Rate           interest.Rate
	MinPayment     money.Money
	DueDay         int
}

type Period struct {
	Loan             string
	Date             *carbon.Carbon
	InterestAccrued  string
	PaymentMade      string
	ResultingBalance string
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

// Make a payment on the loan. Returns remainder, which is non-zero
// when the payment pays off the loan.
func (loan *Loan) PayOnLoan(amount money.Money) (remainder money.Money) {
	// If the minimum doesn't cover interest, sucks to suck. GL;HF
	// if loan.UnpaidInterest.GreaterThanOrEqualTo(amount) {
	// 	loan.UnpaidInterest.Cents = loan.UnpaidInterest.Cents - amount.Cents
	// 	return money.Money{Cents: 0}
	// }

	// How much do we have left after interest?
	centsLeft := amount.Cents - loan.UnpaidInterest.Cents

	if centsLeft > 0 && centsLeft > loan.Principal.Cents {
		remainder := centsLeft - loan.Principal.Cents
		loan.UnpaidInterest.Cents = 0
		loan.Principal.Cents = 0

		return money.Money{Cents: remainder}
	} else if centsLeft > 0 && centsLeft < loan.Principal.Cents {
		loan.UnpaidInterest.Cents = 0
		loan.Principal.Cents = max(loan.Principal.Cents-centsLeft, 0)
		return money.Money{Cents: 0}

	} else {
		loan.UnpaidInterest.Cents += centsLeft
		return money.Money{Cents: 0}
	}

}

func (loan Loan) Equals(other Loan) bool {
	if loan.Name != other.Name {
		return false
	}
	if loan.Principal != other.Principal {
		return false
	}
	if loan.UnpaidInterest != other.UnpaidInterest {
		return false
	}
	if loan.Rate != other.Rate {
		return false
	}
	if loan.MinPayment != other.MinPayment {
		return false
	}
	if loan.DueDay != other.DueDay {
		return false
	}

	return true
}

func (loan Loan) NotEquals(other Loan) bool {
	return !loan.Equals(other)
}
