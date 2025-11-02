package payments

import (
	"fmt"

	"github.com/landoware/debt-deleter/debts"
	"github.com/landoware/debt-deleter/interest"
	"github.com/landoware/debt-deleter/money"
	"github.com/uniplaces/carbon"
)

type State struct {
	InterestAccrued money.Money
	Budget          money.Money
	Loans           []debts.Loan
	Date            *carbon.Carbon
	BestResult      []debts.Loan
}

// Makes payments on the slide of loans contained in state.Loans, according to state.Budget.
// The last loan in the slice gets the remaining budgeted amount.
//
// Returns interestPaid, the total accrued during this permutation, and whether the loans were all
// paidInFull at the end of the function.
// If this permutation accrues more interest paid than bestInterest, the function returns early
// since a prior attempt did better.
func MakePayments(state *State, bestInterest money.Money) (interestPaid money.Money, paidInFull bool) {
	// TODO is fucked
	fmt.Println("making payments\n----------------")
	fmt.Printf("State:\nInterestAccrued: %s\nLoans: %+v\nDate: %s\n\n", state.InterestAccrued.String(), state.Loans, state.Date.DateString())
	// If everything is paid off, return
	if checkPaidOff(state.Loans) {
		return state.InterestAccrued, true
	}

	// Should we even continue? If we're doing worse than our best attempt, nope.
	if state.InterestAccrued.GreaterThan(bestInterest) {
		fmt.Printf("returned because state.InterestAccrued > bestInterest: %s > %s\n", state.InterestAccrued.String(), bestInterest.String())
		return state.InterestAccrued, false
	}

	// Initalize the budgeted amount
	budgetRemaining := state.Budget

	// Filter out Paid Loans
	activeLoans := state.Loans[:0]
	for _, loan := range state.Loans {
		if loan.Principal.GreaterThanZero() {
			activeLoans = append(activeLoans, loan)
		}
	}
	state.Loans = activeLoans

	// Calculate values for each loan and apply the payments
	for i, loan := range state.Loans {
		fmt.Printf("Calculating for %+v\n", loan)

		// Figure out interest
		newInterest := interest.MonthlyInterest(*state.Date, loan.Principal, loan.Rate)
		// Add it to the loan
		loan.UnpaidInterest = loan.UnpaidInterest.Add(newInterest)
		// Add to the total in the state
		state.InterestAccrued = state.InterestAccrued.Add(newInterest)

		fmt.Printf("After Interest accrual: %+v\n", loan)

		// Non-end-of-slice indexes get the minimum payment.
		if i < len(state.Loans)-1 {
			fmt.Println("-- Paying Min Payment --")
			budgetRemaining = budgetRemaining.Subtract(loan.MinPayment)
			remainder := loan.PayOnLoan(loan.MinPayment)
			budgetRemaining = budgetRemaining.Add(remainder)
		} else {
			fmt.Printf("!! PAYING %s !!\n", budgetRemaining.String())
			// Make the extra payment on the last index
			loan.PayOnLoan(budgetRemaining)
		}
		fmt.Printf("After Payment (index %d): %+v\n", i, loan)

		// Persist it to the state
		state.Loans[i] = loan

	}

	// Increment the date
	state.Date = state.Date.AddMonth()

	// Do it all again
	return MakePayments(state, bestInterest)
}

// Are each of the loans in the list fully paid off?
func checkPaidOff(loans []debts.Loan) bool {
	for _, loan := range loans {
		if loan.Principal.GreaterThanZero() {
			return false
		}
	}
	return true
}
