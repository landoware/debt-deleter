package tree

import (
	"fmt"

	"github.com/landoware/debt-deleter/debts"
	"github.com/landoware/debt-deleter/money"
)

const MaxInt = int(^uint(0) >> 1)

type Path struct {
	Keys            []int
	InterestAccrued money.Money
}

type PathCache struct {
	paths []Path
}

func Optimize(
	loans []debts.Loan,
	budget money.Money,
) (
	bestPath Path,
	bestInterest money.Money,
) {
	bestInterest = money.Money{Cents: MaxInt}

	cache := &PathCache{}

	path := Path{}

	cycles := 0
	performCalculations := true
	for performCalculations {
		cycles++
		if cycles > 100 {
			performCalculations = false
		}
		// Break if we're paid in full

		// Break if we've accrued more interest than a better attempt

		// Break if we have a cache match
		if cache.checkMatch(path) {
			fmt.Println("Cache Broken") // Todo, make an exit condition
			performCalculations = false
		}

		// Where oh where shall I put all this money?
		chosen := cycles % 2

		var minPayments money.Money
		for _, loan := range loans {
			if loan.Principal.GreaterThanZero() {
				minPayments = minPayments.Add(loan.MinPayment)
			}
		}
		extraAllocation := budget.Minus(minPayments)

		path.Keys = append(path.Keys, chosen)
		for i, loan := range loans {
			if i == chosen {
				loan.PayOnLoan(extraAllocation)
			} else {
				loan.PayOnLoan(loan.MinPayment)
			}
		}
	}

	cache.paths = append(cache.paths, path)
	path = Path{}

	cycles = 0
	performCalculations = true
	for performCalculations {
		cycles++
		if cycles > 100 {
			performCalculations = false
		}
		// Break if we're paid in full

		// Break if we have a cache match
		if cache.checkMatch(path) {
			fmt.Println("Cache Broken") // Todo, make an exit condition
			performCalculations = false
		}

		// Where oh where shall I put all this money?
		chosen := cycles % 2

		var minPayments money.Money
		for _, loan := range loans {
			if loan.Principal.GreaterThanZero() {
				minPayments = minPayments.Add(loan.MinPayment)
			}
		}
		extraAllocation := budget.Minus(minPayments)

		path.Keys = append(path.Keys, chosen)
		for i, loan := range loans {
			if i == chosen {
				loan.PayOnLoan(extraAllocation)
			} else {
				loan.PayOnLoan(loan.MinPayment)
			}
		}
	}
	bestPath = path

	return bestPath, bestInterest
}

func (cache *PathCache) checkMatch(needle Path) bool {
	for _, path := range cache.paths {
		if path.equals(needle) {
			return true
		}
	}
	return false
}

func (p Path) equals(other Path) bool {
	for i, loanIndex := range p.Keys {
		if loanIndex != other.Keys[i] {
			return false
		}
	}
	return true
}
