package payments

import (
	"fmt"
	"math/rand"

	"github.com/landoware/debt-deleter/debts"
	"github.com/landoware/debt-deleter/interest"
	"github.com/landoware/debt-deleter/money"
	"github.com/uniplaces/carbon"
)

type State struct {
	interestAccrued money.Money
	budget          money.Money
	loans           []debts.Loan
	date            *carbon.Carbon
	paidExtraOn     []int
	bestResult      []int
}

const MaxInt = int(^uint(0) >> 1)

// Given a budget and a slice of loans, try to minimize the amount of interest
// that will accrue until the loans are paid off.
// The allocations result will be the index in loans which the higher-than-minimum
// payment was made to.
//
// A sucessful optimization will return err = nil. However, if the maximum attempts
// were exceeded, a usuable value may be in allocations.
func OptimizeLoans(budget money.Money, loans []debts.Loan) (allocations []int, totalInterestAccrued money.Money, attempts int, err error) {

	state := State{
		interestAccrued: money.NewMoney(0, 0),
		budget:          budget,
		loans:           loans,
		date:            carbon.Now(),
	}

	bestInterest := money.Money{Cents: MaxInt}

	attempts = 0

	for {
		attempts++

		if attempts > 100 {
			return state.bestResult, bestInterest, attempts, fmt.Errorf("Stopped after %d attempts", attempts)
		}

		totalInterestAccrued, paidInFull := MakePayments(&state, bestInterest)
		// totalInterestAccrued, _ := MakePayments(&state, bestInterest)

		if bestInterest == totalInterestAccrued {
			return state.bestResult, totalInterestAccrued, attempts, nil
		}

		if paidInFull && totalInterestAccrued.LessThan(bestInterest) {
			state.bestResult = state.paidExtraOn
			bestInterest = totalInterestAccrued
		}

	}

}

// Recusrive method to allocate a budgeted amount across the list of loans in the most efficient way.
// Should be able to take in any state and figure out "from this point, what's the best allocation
// to minimize interest paid?"
//
// Returns itnerestPaid, the total accrued during this payment attempt, and whether the loans were all
// paidInFull at the end of the function.
func MakePayments(state *State, bestInterest money.Money) (interestPaid money.Money, paidInFull bool) {
	// If everything is paid off, return
	if checkPaidOff(state.loans) {
		return state.interestAccrued, true
	}

	// Should we even continue? If we're doing worse than our best attempt, nope.
	if state.interestAccrued.GreaterThan(bestInterest) {
		return state.interestAccrued, false
	}

	// Pick the loan to make the largest payment to at random.
	chosenIndex := rand.Intn(len(state.loans))
	budgetRemaining := state.budget

	for i, loan := range state.loans {

		// Figure out interest
		newInterest := interest.MonthlyInterest(*state.date, loan.Principal, loan.Rate)
		loan.UnpaidInterest = loan.UnpaidInterest.Add(newInterest)

		unpaidInterest := loan.UnpaidInterest
		state.interestAccrued = state.interestAccrued.Add(unpaidInterest)
		state.paidExtraOn = append(state.paidExtraOn, chosenIndex)

		if i != chosenIndex {
			budgetRemaining = budgetRemaining.Subtract(loan.MinPayment)
			loan.PayOnLoan(loan.MinPayment)
		}
	}

	// Finally, pay the remaining budget on the chosen loan
	state.loans[chosenIndex].PayOnLoan(budgetRemaining)

	// Increment the date
	state.date.AddMonth()

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
