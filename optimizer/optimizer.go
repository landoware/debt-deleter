package optimizer

import (
	"github.com/landoware/debt-deleter/debts"
	"github.com/landoware/debt-deleter/money"
	"github.com/landoware/debt-deleter/payments"
	"github.com/uniplaces/carbon"
)

const MaxInt = int(^uint(0) >> 1)

// Given a budget and a slice of loans, try to minimize the amount of interest
// that will accrue until the loans are paid off.
// The allocations result will be the index in loans which the higher-than-minimum
// payment was made to.
//
// A sucessful optimization will return err = nil. However, if the maximum attempts
// were exceeded, a usuable value may be in allocations.
func Optimize(loans []debts.Loan, budget money.Money) (loansOrderedToPay []debts.Loan, interestAccrued money.Money) {

	state := payments.State{
		InterestAccrued: money.NewMoney(0, 0),
		Budget:          budget,
		Loans:           loans,
		Date:            carbon.Now(),
	}

	bestInterest := money.Money{Cents: MaxInt}

	length := len(loans)

	for range length {
		// Make a copy of the original state of the loans
		unalteredLoans := deepCopy(loans)
		// Make payments on each loan, paying the minimum on everything except the last loan in the slice
		totalInterestAccrued, paidInFull := payments.MakePayments(&state, bestInterest)

		// How'd we do?
		if paidInFull && totalInterestAccrued.LessThan(bestInterest) {
			state.BestResult = deepCopy(unalteredLoans)
			bestInterest = totalInterestAccrued
		}

		// Reorder the loans slice and set up for the next run
		loans = HeapsAlgorithm(length, unalteredLoans)
		state.Loans = loans
		state.InterestAccrued.Cents = 0
		state.Date = carbon.Now()
	}

	return state.BestResult, bestInterest

}

// Reorders the loans provided. When called the number of times corresponding to its length,
// this algorithm is guaranteed to have returned every possible permutation.
// https://en.wikipedia.org/wiki/Heap%27s_algorithm
func HeapsAlgorithm(length int, loans []debts.Loan) []debts.Loan {
	if length > 1 {
		// loans := HeapsAlgorithm(length, loans)
		for i := range length - 1 {
			if length%2 == 0 {
				loans[i], loans[length-1] = loans[length-1], loans[i]
			} else {
				loans[0], loans[length-1] = loans[length-1], loans[0]
			}
			return HeapsAlgorithm(length-1, loans)
		}
	}
	return loans
}

func deepCopy(loans []debts.Loan) (copy []debts.Loan) {
	copy = make([]debts.Loan, len(loans))

	for i, loan := range loans {
		copy[i].Name = loan.Name
		copy[i].Principal = loan.Principal
		copy[i].UnpaidInterest = loan.UnpaidInterest
		copy[i].Rate = loan.Rate
		copy[i].MinPayment = loan.MinPayment
		copy[i].DueDay = loan.DueDay
	}

	return copy
}
