package main

import (
	"errors"
	"fmt"
	"github.com/uniplaces/carbon"
	"math"
	"time"
)

type Loan struct {
	name        string
	balance     float64
	rate        float64
	min_payment float64
	due_day     int
}

func newLoan(name string, balance, rate, min_payment float64, due_day int) Loan {
	return Loan{
		name:        name,
		balance:     balance,
		rate:        rate,
		min_payment: min_payment,
		due_day:     due_day,
	}
}

func main() {
	loanA := newLoan("Nelnet AA", 5000.00, 0.05, 150.0, 18)

	fmt.Printf("Amortizing %.2f with payments of $%.2f\n\n", loanA.balance, loanA.min_payment)

	if loanAPayoff, loanASchedule, err := getNumberOfPayments(loanA.balance, loanA.rate, 150, loanA.due_day); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s will pay off in %d payments", loanA.name, loanAPayoff)
		fmt.Printf("\n\n%v", loanASchedule)
	}
}

// @todo reutrn an array of the balances
func getNumberOfPayments(balance, rate, payment float64, paymentDay int) (int, []float64, error) {
	var schedule []float64

	if payment >= balance {
		return 1, schedule, nil
	}

	if balance <= 0 {
		return 0, schedule, nil
	}

	payments := 0

	startDate, carbonErr := carbon.CreateFromDate(carbon.Now().Year(), carbon.Now().Month(), paymentDay, time.Local.String())
	if carbonErr != nil {
		return 0, schedule, carbonErr
	}

	if startDate.Before(time.Now()) {
		startDate.AddDate(0, 1, 0)
	}

	for balance > 0 {
		if payments > 1000 {
			return 0, schedule, errors.New("Failed to amortize.\n")
		}
		payments++
		balance = balance + getMonthlyInterest(*startDate, balance, rate) - payment
		schedule = append(schedule, math.Max(balance, 0))

	}
	return payments, schedule, nil
}

func getMonthlyInterest(startDate carbon.Carbon, balance, rate float64) float64 {
	endDate := startDate.AddMonth()

	days := endDate.DiffInDays(&startDate, true)

	return balance * rate * float64(days) / 365.25

}
