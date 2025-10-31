package payments_test

import (
	"testing"

	"github.com/landoware/debt-deleter/debts"
	"github.com/landoware/debt-deleter/interest"
	"github.com/landoware/debt-deleter/money"
	"github.com/landoware/debt-deleter/payments"
)

func OptimizeSingleLoan(t *testing.T) {
	rate := interest.NewRateFromParts(5, 0)
	var loans []debts.Loan
	loan := debts.NewLoan("test", money.NewMoney(1000, 0), rate, money.NewMoney(50, 0), 1, money.NewMoney(0, 0))
	loans = append(loans, loan)
	budget := money.NewMoney(75, 0)

	payments.OptimizeLoans(budget, loans)
}
