package payments

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/landoware/debt-deleter/debts"
	"github.com/landoware/debt-deleter/interest"
	"github.com/landoware/debt-deleter/money"
	"github.com/uniplaces/carbon"
)

type State struct {
	interestAccrued     money.Money
	bestInterestAccrued money.Money
	budget              money.Money
	loans               []debts.Loan
	date                *carbon.Carbon
}

const MaxInt = int(^uint(0) >> 1)

func OptimizeLoans(budget money.Money, loans []debts.Loan) ([]debts.Loan, error) {

	state := State{
		interestAccrued: money.NewMoney(0, 0),
		budget:          budget,
		loans:           loans,
		date:            carbon.Now(),
	}

	var bestResult []debts.Loan
	bestInterest := money.Money{Cents: MaxInt}

	attempts := 0

	paidLoans, interestPaid := MakePayments(&state, bestInterest)
	for {
		attempts++
		if attempts > 1000 {
			return bestResult, fmt.Errorf("Stopped after %d attempts", attempts)
		}

		if paidLoans != nil && interestPaid.LessThan(bestInterest) {
			bestResult = paidLoans
			bestInterest = interestPaid
		}

		MakePayments(&state, bestInterest)
	}

	return bestResult, nil
}

// Recusrive method to allocate a budgeted amount across the list of loans in the most efficient way.
// Should be able to take in any state and figure out "from this point, what's the best allocation
// to minimize interest paid?"
func MakePayments(state *State, bestInterest money.Money) ([]debts.Loan, money.Money) {
	// If everything is paid off, return
	if checkPaidOff(state.loans) {
		return state.loans, state.interestAccrued
	}

	// Should we even continue? If we're doing worse than our best attempt, nope.
	if state.interestAccrued.GreaterThan(state.bestInterestAccrued) {
		return nil, state.interestAccrued
	}

	// Pick the loan to make the largest payment to at random.
	// Fuck it.
	chosenIndex := rand.Intn(len(state.loans))
	budgetRemaining := state.budget

	for i, loan := range state.loans {

		// Figure out interest
		newInterest := interest.MonthlyInterest(*state.date, loan.Principal, loan.Rate)
		loan.UnpaidInterest = loan.UnpaidInterest.Add(newInterest)

		unpaidInterest := loan.UnpaidInterest

		if i != chosenIndex {
			budgetRemaining = budgetRemaining.Subtract(loan.MinPayment)
			loan.PayOnLoan(loan.MinPayment)

			payment := debts.Payment{Date: *state.date, Amount: loan.MinPayment}

			loan.Schedule = append(loan.Schedule, payment)

			// Set the running interest paid in the state.
			if loan.MinPayment.LessThanOrEqualTo(unpaidInterest) {
				state.interestAccrued = state.interestAccrued.Add(unpaidInterest)
			} else {
				state.interestAccrued = state.interestAccrued.Add(loan.MinPayment)
			}
		}
	}

	// Finally, pay the remaining budget on the chosen loan
	state.loans[chosenIndex].PayOnLoan(budgetRemaining)

	// Increment the date
	state.date.AddMonth()

	// Do it all again
	return MakePayments(state, bestInterest)
}

func checkPaidOff(loans []debts.Loan) bool {
	for _, loan := range loans {
		if loan.Principal.GreaterThanZero() {
			return false
		}
	}
	return true
}
