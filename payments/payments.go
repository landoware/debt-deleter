package payments

import (
	"github.com/landoware/debt-deleter/debts"
	"github.com/landoware/debt-deleter/interest"
	"github.com/landoware/debt-deleter/money"
	"github.com/uniplaces/carbon"
	"math/rand"
)

type State struct {
	interestPaid     money.Money
	bestInterestPaid money.Money
	budget           money.Money
	loans            []debts.Loan
	date             *carbon.Carbon
}

const MaxInt = int(^uint(0) >> 1)

func OptimizeLoans(budget money.Money, loans []debts.Loan) {

	state := State{
		interestPaid:     money.NewMoney(0, 0),
		bestInterestPaid: money.Money{Cents: MaxInt},
		budget:           budget,
		loans:            loans,
		date:             carbon.Now(),
	}

	paidLoans := MakePayments(&state)
	for {
		if paidLoans != nil {
			break
		}
	}

}

// Recusrive method to allocate a budgeted amount across the list of loans in the most efficient way.
// Should be able to take in any state and figure out "from this point, what's the best allocation
// to minimize interest paid?"
func MakePayments(state *State) []debts.Loan {
	// If everything is paid off, return
	if checkPaidOff(state.loans) {
		if state.interestPaid.LessThan(state.bestInterestPaid) {
			state.bestInterestPaid = state.interestPaid
		}
		return state.loans
	}

	// Should we even continue? If we're doing worse than our best attempt, nope.
	if state.interestPaid.GreaterThan(state.bestInterestPaid) {
		return nil
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
				state.interestPaid = state.interestPaid.Add(unpaidInterest)
			} else {
				state.interestPaid = state.interestPaid.Add(loan.MinPayment)
			}
		}
	}

	state.loans[chosenIndex].PayOnLoan(budgetRemaining)

	// Add a month to the date
	// do it all again
	return MakePayments(state)
}

func checkPaidOff(loans []debts.Loan) bool {
	for _, loan := range loans {
		if loan.Principal.GreaterThanZero() {
			return false
		}
	}
	return true
}
