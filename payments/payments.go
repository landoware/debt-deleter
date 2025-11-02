package payments

import (
	"time"

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
	Schedule        []debts.Period
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
	// If everything is paid off, return
	if checkPaidOff(state.Loans) {
		return state.InterestAccrued, true
	}

	// Should we even continue? If we're doing worse than our best attempt, nope.
	if state.InterestAccrued.GreaterThan(bestInterest) {
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

		dueDate := getDueDateFromDate(loan, state.Date)

		// Figure out interest
		newInterest := interest.MonthlyInterest(*dueDate, loan.Principal, loan.Rate)
		// Add it to the loan
		loan.UnpaidInterest = loan.UnpaidInterest.Add(newInterest)
		// Add to the total in the state
		state.InterestAccrued = state.InterestAccrued.Add(newInterest)

		period := debts.Period{
			Loan:            loan.Name,
			Date:            dueDate,
			InterestAccrued: newInterest.String(),
		}

		// Non-end-of-slice indexes get the minimum payment.
		if i < len(state.Loans)-1 {
			period.PaymentMade = loan.MinPayment.String()
			budgetRemaining = budgetRemaining.Subtract(loan.MinPayment)
			remainder := loan.PayOnLoan(loan.MinPayment)
			budgetRemaining = budgetRemaining.Add(remainder)
		} else {
			// Make the extra payment on the last index
			period.PaymentMade = budgetRemaining.String()
			loan.PayOnLoan(budgetRemaining)
		}

		// Finish up the schedule's data
		period.ResultingBalance = loan.Principal.String()
		state.Schedule = append(state.Schedule, period)

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

func getDueDateFromDate(loan debts.Loan, date *carbon.Carbon) *carbon.Carbon {
	dueDate, carbonErr := carbon.CreateFromDate(date.Year(), date.Month(), loan.DueDay, time.Local.String())
	if carbonErr != nil {
		return date
	}
	return dueDate
}
