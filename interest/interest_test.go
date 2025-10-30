package interest_test

import (
	"github.com/landoware/debt-deleter/interest"
	"github.com/landoware/debt-deleter/money"
	"github.com/uniplaces/carbon"
	"testing"
	"time"
)

func TestDailyInterest(t *testing.T) {
	balances := []money.Money{money.NewMoney(5000, 0), money.NewMoney(100, 0), money.NewMoney(1000, 0)}
	rates := []interest.Rate{interest.NewRateFromParts(5, 25), interest.NewRateFromParts(10, 125), interest.NewRateFromParts(5, 0)}
	expectedCentsInterest := []int{71, 2, 13}

	for i := range len(balances) {
		dailyInterest := interest.DailyInterest(rates[i], balances[i])

		if dailyInterest.Cents != expectedCentsInterest[i] {
			t.Errorf("Expected %d, got %d", expectedCentsInterest[i], dailyInterest.Cents)
		}
	}
}

func TestMonthlyInterest(t *testing.T) {
	balances := []money.Money{money.NewMoney(5000, 0), money.NewMoney(100, 0), money.NewMoney(1000, 0)}
	rates := []interest.Rate{interest.NewRateFromParts(5, 25), interest.NewRateFromParts(10, 125), interest.NewRateFromParts(5, 0)}
	expectedInterest := []money.Money{money.NewMoney(22, 1), money.NewMoney(0, 62), money.NewMoney(4, 3)}

	date, carbonErr := carbon.CreateFromDate(2025, 1, 1, time.Local.String())
	if carbonErr != nil {
		t.Error(carbonErr)
	}

	for i := range len(balances) {
		result := interest.MonthlyInterest(*date, balances[i], rates[i])

		if result != expectedInterest[i] {
			t.Errorf("Expected %s, got %s", result, expectedInterest[i])
		}
	}
}
