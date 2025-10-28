package money_test

import (
	"github.com/landoware/debt-deleter/money"
	"testing"
)

func TestNewMoney(t *testing.T) {
	m := money.NewMoney(10, 50)
	if m == nil {
		t.Error("Failed to create new money")
	}
}

func TestDollars(t *testing.T) {
	expected := 10
	tenDollars := money.NewMoney(expected, 99)

	result := tenDollars.Dollars()

	if expected != result {
		t.Errorf("Expected %d, got %d", expected, result)
	}

}

func TestOnlyCents(t *testing.T) {
	expected := 99
	ninetyNineCents := money.NewMoney(0, expected)

	result := ninetyNineCents.OnlyCents()

	if expected != result {
		t.Errorf("Expected %d, got %d", expected, result)
	}

}

// func TestCents(t *testing.T) {
// }
//
// func TestAdd(t *testing.T) {
// }
//
// func TestSubtract(t *testing.T) {
// }
