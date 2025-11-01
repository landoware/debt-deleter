package payments_test

import (
	"testing"

	"github.com/landoware/debt-deleter/debts"
	"github.com/landoware/debt-deleter/interest"
	"github.com/landoware/debt-deleter/money"
	"github.com/landoware/debt-deleter/payments"
)

// func TestOptimizeSingleLoan(t *testing.T) {
// 	rate := interest.NewRateFromParts(5, 0)
// 	var loans []debts.Loan
// 	loan := debts.NewLoan("test", money.NewMoney(1000, 0), rate, money.NewMoney(50, 0), 1, money.NewMoney(0, 0))
// 	loans = append(loans, loan)
// 	budget := money.NewMoney(75, 0)
//
// 	result, interest, attempts, err := payments.OptimizeLoans(budget, loans)
//
// 	if err != nil {
// 		t.Logf("Result after %d attempts: %s paid\n%v", attempts, interest.String(), result)
// 	}
// }

func TestOptimizeTwoLoans(t *testing.T) {
	rate := interest.NewRateFromParts(5, 0)
	var loans []debts.Loan

	loanA := debts.NewLoan("test", money.NewMoney(1000, 0), rate, money.NewMoney(50, 0), 1, money.NewMoney(0, 0))
	loans = append(loans, loanA)

	loanB := debts.NewLoan("test", money.NewMoney(1500, 0), rate, money.NewMoney(50, 0), 1, money.NewMoney(0, 0))
	loans = append(loans, loanB)

	budget := money.NewMoney(75, 0)
	// result, interest, attempts, err := payments.OptimizeLoans(budget, loans)
	result, interest, attempts, _ := payments.OptimizeLoans(budget, loans)

	// if err != nil {
	t.Errorf("Result after %d attempts: %s paid\n%v", attempts, interest.String(), result)
	// }
}
