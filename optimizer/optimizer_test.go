package optimizer_test

import (
	"testing"

	"github.com/landoware/debt-deleter/debts"
	"github.com/landoware/debt-deleter/interest"
	"github.com/landoware/debt-deleter/money"
	"github.com/landoware/debt-deleter/optimizer"
)

func TestHeapsAlgrorith(t *testing.T) {
	var loans []debts.Loan
	rate := interest.NewRateFromParts(5, 0)

	loanA := debts.NewLoan("Loan 1", money.NewMoney(1000, 0), rate, money.NewMoney(50, 0), 1, money.NewMoney(0, 0))
	loans = append(loans, loanA)

	loanB := debts.NewLoan("Loan 2", money.NewMoney(1500, 0), rate, money.NewMoney(50, 0), 1, money.NewMoney(0, 0))
	loans = append(loans, loanB)

	loanC := debts.NewLoan("Loan 3", money.NewMoney(1500, 0), rate, money.NewMoney(50, 0), 1, money.NewMoney(0, 0))
	loans = append(loans, loanC)

	result := optimizer.HeapsAlgorithm(len(loans), loans)

	if result[0].Name != loanB.Name {
		t.Errorf("Expected %s in index 0, got %s\n%v", loanB.Name, result[0].Name, result)
	}
	if result[1].Name != loanC.Name {
		t.Errorf("Expected %s in index 1, got %s\n%v", loanC.Name, result[1].Name, result)
	}
	if result[2].Name != loanA.Name {
		t.Errorf("Expected %s in index 2, got %s\n%v", loanA.Name, result[2].Name, result)
	}
}

func TestOptimizeSingleLoan(t *testing.T) {
	rate := interest.NewRateFromParts(5, 0)
	var loans []debts.Loan
	loan := debts.NewLoan("test", money.NewMoney(1000, 0), rate, money.NewMoney(50, 0), 1, money.NewMoney(0, 0))
	loans = append(loans, loan)
	budget := money.NewMoney(75, 0)

	var expected []debts.Loan
	expected = append(expected, loan)

	result, _ := optimizer.Optimize(loans, budget)

	for i := range result {
		if result[i].NotEquals(expected[i]) {
			t.Errorf("Expected '%s' to be in index %d, '%s' found", expected[i].Name, i, result[i].Name)
		}
	}
}

func TestOptimizeTwoLoans(t *testing.T) {
	var loans []debts.Loan

	rateA := interest.NewRateFromParts(5, 0)
	loanA := debts.NewLoan("Loan A", money.NewMoney(1000, 0), rateA, money.NewMoney(50, 0), 1, money.NewMoney(0, 0))
	loans = append(loans, loanA)

	rateB := interest.NewRateFromParts(10, 0)
	loanB := debts.NewLoan("Loan B", money.NewMoney(5000, 0), rateB, money.NewMoney(50, 0), 1, money.NewMoney(0, 0))
	loans = append(loans, loanB)

	var expected []debts.Loan
	expected = append(expected, loanB)
	expected = append(expected, loanA)

	budget := money.NewMoney(150, 0)

	result, _ := optimizer.Optimize(loans, budget)

	for i := range result {
		if result[i].NotEquals(expected[i]) {
			t.Logf("\nresult %+v\nexpected %+v\n\n", result, expected)
			t.Errorf("Expected '%s' with balance %s to be in index %d, '%s' with %s found", expected[i].Name, expected[i].Principal.String(), i, result[i].Name, result[i].Principal.String())
		}
	}
}
