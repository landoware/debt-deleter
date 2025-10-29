package interest

import (
	"github.com/landoware/debt-deleter/money"
	"math"
	"testing"
)

type Rate struct {
	whole_part      int
	fractional_part int
}

func NewRateFromParts(whole_part, fractional_part int) Rate {
	return Rate{
		whole_part:      whole_part,
		fractional_part: fractional_part,
	}
}

func (rate Rate) convertForCalculation() (int, int) {
	exponent := getIntegerLength(rate.fractional_part)
	factor := int(math.Pow10(exponent))

	integerRate := rate.whole_part*factor + rate.fractional_part

	return integerRate, factor
}

func getIntegerLength(fractional_part int) int {
	if fractional_part < 10 {
		// Default to two decimal places so that we can use 365.25 as a day basis
		return 2
	}
	return int(math.Floor(math.Log10(float64(fractional_part)))) + 1
}

func DailyInterest(rate Rate, balance money.Money) money.Money {
	integerRate, factor := rate.convertForCalculation()

	// Cents are 100x larger than percentages, so reduce the order of magnitude after scaling them up by the interest factor.
	cents := balance.Cents * factor / 100

	//
	dayBasis := 36525 * factor / 100

	result := cents * integerRate / dayBasis / factor

	return money.NewMoney(0, result)
}

// Unit tests for unexported functions

func TestConvertForCalculation(t *testing.T) {
	rates := []Rate{NewRateFromParts(5, 0), NewRateFromParts(5, 25)}
	expectedResults := [][]int{{5, 1}, {525, 2}}

	for i, expectedResult := range expectedResults {
		integerRate, factor := rates[i].convertForCalculation()
		if integerRate != expectedResult[0] && factor != expectedResult[1] {
			t.Errorf("Expected Integer Rate %d and Factor %d, got rate %d with factor %d", integerRate, factor, expectedResult[0], expectedResult[1])
		}
		if integerRate != expectedResult[0] {
			t.Errorf("Expected Integer Rate %d, got %d", integerRate, expectedResult[0])
		}
		if factor != expectedResult[1] {
			t.Errorf("Expected Factor %d, got %d", factor, expectedResult[1])
		}
	}
}

func TestGetIntegerLength(t *testing.T) {
	tests := [][]int{{0, 1}, {1, 1}, {10, 2}, {99, 2}, {100, 3}, {1000, 4}}

	for _, test := range tests {
		length := getIntegerLength(test[0])
		if length == test[1] {
			t.Errorf("Expected length %d, got %d", test[1], length)
		}
	}
}
