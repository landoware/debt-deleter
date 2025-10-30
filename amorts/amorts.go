package main

import (
	"errors"
	"github.com/landoware/debt-deleter/interest"
	"github.com/landoware/debt-deleter/money"
	"github.com/uniplaces/carbon"
	"time"
)

// @todo reutrn an array of the balances
func getNumberOfPayments(balance money.Money, rate interest.Rate, payment money.Money, paymentDay int) (int, []money.Money, error) {
	var schedule []money.Money

	if payment.GreaterThanOrEqualTo(balance) {
		return 1, schedule, nil
	}

	if balance.LessThanOrEqualToZero() {
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

	for balance.GreaterThanZero() {
		if payments > 1000 {
			return payments, schedule, errors.New("Failed to amortize.\n")
		}
		payments++

		// todo figure out how to properly chain .Add().Subtract()
		balance = balance.Add(interest.MonthlyInterest(*startDate, balance, rate)).Subtract(payment)
		schedule = append(schedule, balance)
	}
	return payments, schedule, nil
}
