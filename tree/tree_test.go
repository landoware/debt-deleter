package tree_test

import (
	"testing"

	"github.com/landoware/debt-deleter/debts"
	"github.com/landoware/debt-deleter/interest"
	"github.com/landoware/debt-deleter/money"
	"github.com/landoware/debt-deleter/tree"
)

func TestOptimize(t *testing.T) {
	loans := []debts.Loan{
		debts.NewLoan("Test A", money.NewMoney(1000, 0), interest.NewRateFromParts(5, 0), money.NewMoney(50, 0), 1, money.NewMoney(0, 0)),
		debts.NewLoan("Test B", money.NewMoney(5000, 0), interest.NewRateFromParts(5, 0), money.NewMoney(50, 0), 1, money.NewMoney(0, 0)),
	}
	budget := money.NewMoney(250, 0)

	bestPath, bestInterest := tree.Optimize(loans, budget)
	t.Log(bestPath)
	t.Log(bestInterest)

	t.FailNow()
}
