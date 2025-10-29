package interest_test

import (
	"github.com/landoware/debt-deleter/interest"
	"github.com/landoware/debt-deleter/money"
	"testing"
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
