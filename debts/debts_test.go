package debts_test

import (
	"testing"

	"github.com/landoware/debt-deleter/debts"
	"github.com/landoware/debt-deleter/interest"
	"github.com/landoware/debt-deleter/money"
)

func TestPayOnLoanWithoutInterest(t *testing.T) {
	loan := debts.Loan{
		Name:           "test",
		Principal:      money.NewMoney(100, 0),
		UnpaidInterest: money.NewMoney(0, 0),
		Rate:           interest.NewRateFromParts(5, 0),
		MinPayment:     money.NewMoney(50, 0),
		DueDay:         1,
	}

	loan.PayOnLoan(loan.MinPayment)

	expected := money.NewMoney(50, 0)

	if loan.Principal.NotEquals(expected) {
		t.Errorf("Expected balance after %s payment to be %s, %s returned", loan.MinPayment, expected, loan.Principal)
	}

}
