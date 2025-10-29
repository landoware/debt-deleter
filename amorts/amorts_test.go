package main

import (
	"github.com/landoware/debt-deleter/interest"
	"github.com/landoware/debt-deleter/money"
	"testing"
)

func TestGetNumberOfPayments(t *testing.T) {
	expected := 79

	balance := money.NewMoney(10000, 0)
	rate := interest.NewRateFromParts(5, 0)
	payment := money.NewMoney(150, 0)

	// func getNumberOfPayments(balance money.Money, rate interest.Rate, payment money.Money, paymentDay int) (int, []money.Money, error) {
	result, _, err := getNumberOfPayments(balance, rate, payment, 15)

	if err != nil {
		t.Error(err)
	}

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
