package main

import (
	"fmt"

	"github.com/landoware/debt-deleter/debts"
	"github.com/landoware/debt-deleter/interest"
	"github.com/landoware/debt-deleter/money"
	"github.com/landoware/debt-deleter/optimizer"
)

func main() {

	budget := money.NewMoney(500, 0)

	var loans []debts.Loan

	loanA := debts.NewLoan(
		"Example 1",
		money.NewMoney(5000, 93),         // Prin
		interest.NewRateFromParts(4, 53), // Rate
		money.NewMoney(56, 63),           // Min Payment
		16,                               // Due day
		money.NewMoney(17, 65),           // Unpaid Int
	)
	loans = append(loans, loanA)

	loanB := debts.NewLoan(
		"Example 2",
		money.NewMoney(7439, 88),         // Prin
		interest.NewRateFromParts(2, 75), // Rate
		money.NewMoney(51, 30),           // Min Payment
		16,                               // Due day
		money.NewMoney(43, 68),           // Unpaid Int
	)
	loans = append(loans, loanB)
	//
	//

	orderToPay, totalInterest, schedule := optimizer.Optimize(loans, budget)

	for _, loan := range orderToPay {
		fmt.Printf("\n%s, ", loan.Name)
	}
	fmt.Printf("\nTotal Interest:, %s", totalInterest.String())
	fmt.Printf("\nFinal Payment:, %s", schedule[len(schedule)-1].Date.DateString())

}
