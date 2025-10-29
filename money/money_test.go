package money_test

import (
	"github.com/landoware/debt-deleter/money"
	"testing"
)

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

func TestAdd(t *testing.T) {
	a := money.NewMoney(10, 50)
	b := money.NewMoney(1, 25)

	expected := money.NewMoney(11, 75)
	result := a.Add(b)

	if result.NotEquals(expected) {
		t.Errorf("Expected %s, got %s", expected.String(), result.String())
	}
}

func TestSubtract(t *testing.T) {
	a := money.NewMoney(10, 50)
	b := money.NewMoney(1, 25)

	expected := money.NewMoney(9, 25)
	result := a.Subtract(b)

	if result.NotEquals(expected) {
		t.Errorf("Expected %s, got %s", expected.String(), result.String())
	}
}

func TestLessThan(t *testing.T) {
	lesser := money.NewMoney(0, 20)
	greater := money.NewMoney(20, 0)

	lesserResult := lesser.LessThan(greater)
	greaterResult := greater.LessThan(lesser)

	equalResult := lesser.LessThan(lesser)

	if lesserResult != true {
		t.Errorf("Expected %s < %s, got %t", lesser.String(), greater.String(), lesserResult)
	}

	if greaterResult != false {
		t.Errorf("Expected %s not less than %s, got %t", greater.String(), lesser.String(), greaterResult)
	}

	if equalResult != false {
		t.Errorf("Expected %s != %s, got %t", lesser.String(), greater.String(), equalResult)
	}

}

func TestLessThanOrEqualTo(t *testing.T) {
	lesser := money.NewMoney(0, 20)
	greater := money.NewMoney(20, 0)

	lesserResult := lesser.LessThanOrEqualTo(greater)
	greaterResult := greater.LessThanOrEqualTo(lesser)

	equalResult := lesser.LessThanOrEqualTo(lesser)

	if lesserResult != true {
		t.Errorf("Expected %s < %s, got %t", lesser.String(), greater.String(), lesserResult)
	}

	if greaterResult != false {
		t.Errorf("Expected %s not less than %s, got %t", greater.String(), lesser.String(), greaterResult)
	}

	if equalResult != true {
		t.Errorf("Expected %s = %s, got %t", lesser.String(), lesser.String(), equalResult)
	}

}

func TestGreaterThan(t *testing.T) {
	lesser := money.NewMoney(0, 20)
	greater := money.NewMoney(20, 0)

	lesserResult := lesser.GreaterThan(greater)
	greaterResult := greater.GreaterThan(lesser)

	equalResult := lesser.GreaterThan(lesser)

	if lesserResult != false {
		t.Errorf("Expected %s not greater than %s, got %t", lesser.String(), greater.String(), lesserResult)
	}

	if greaterResult != true {
		t.Errorf("Expected %s > %s, got %t", greater.String(), lesser.String(), greaterResult)
	}

	if equalResult != false {
		t.Errorf("Expected %s != %s, got %t", lesser.String(), greater.String(), equalResult)
	}

}

func TestGreaterThanOrEqualTo(t *testing.T) {
	lesser := money.NewMoney(0, 20)
	greater := money.NewMoney(20, 0)

	lesserResult := lesser.GreaterThanOrEqualTo(greater)
	greaterResult := greater.GreaterThanOrEqualTo(lesser)

	equalResult := lesser.GreaterThanOrEqualTo(lesser)

	if lesserResult != false {
		t.Errorf("Expected %s < %s, got %t", lesser.String(), greater.String(), lesserResult)
	}

	if greaterResult != true {
		t.Errorf("Expected %s not greater than %s, got %t", greater.String(), lesser.String(), greaterResult)
	}

	if equalResult != true {
		t.Errorf("Expected %s = %s, got %t", lesser.String(), lesser.String(), equalResult)
	}

}

func TestGreaterThanOrEqualToZero(t *testing.T) {
	zero := money.NewMoney(0, 0)
	positive := money.NewMoney(20, 0)
	negative := money.NewMoney(-20, 0)

	if zero.GreaterThanOrEqualToZero() != true {
		t.Errorf("Expected %s to be >= 0, got %t", zero.String(), zero.GreaterThanOrEqualToZero())
	}

	if positive.GreaterThanOrEqualToZero() != true {
		t.Errorf("Expected %s to be >= 0, got %t", positive.String(), positive.GreaterThanOrEqualToZero())
	}

	if negative.GreaterThanOrEqualToZero() != false {
		t.Errorf("Expected %s to not be greater than or equal to 0, got %t", negative.String(), negative.GreaterThanOrEqualToZero())
	}

}

func TestGreaterThanZero(t *testing.T) {
	zero := money.NewMoney(0, 0)
	positive := money.NewMoney(20, 0)
	negative := money.NewMoney(-20, 0)

	if zero.GreaterThanZero() != false {
		t.Errorf("Expected %s to be not equal to 0, got %t", zero.String(), zero.GreaterThanZero())
	}

	if positive.GreaterThanZero() != true {
		t.Errorf("Expected %s to be > 0, got %t", positive.String(), positive.GreaterThanZero())
	}

	if negative.GreaterThanZero() != false {
		t.Errorf("Expected %s to not be greater than 0, got %t", negative.String(), negative.GreaterThanZero())
	}
}

func TestLesserThanOrEqualToZero(t *testing.T) {
	zero := money.NewMoney(0, 0)
	positive := money.NewMoney(20, 0)
	negative := money.NewMoney(-20, 0)

	if zero.LessThanOrEqualToZero() != true {
		t.Errorf("Expected %s to be <= 0, got %t", zero.String(), zero.LessThanOrEqualToZero())
	}

	if positive.LessThanOrEqualToZero() != false {
		t.Errorf("Expected %s to be less than or equal to 0, got %t", positive.String(), positive.LessThanOrEqualToZero())
	}

	if negative.LessThanOrEqualToZero() != true {
		t.Errorf("Expected %s to be greater than or equal to 0, got %t", negative.String(), negative.LessThanOrEqualToZero())
	}

}

func TestLesserThanZero(t *testing.T) {
	zero := money.NewMoney(0, 0)
	positive := money.NewMoney(20, 0)
	negative := money.NewMoney(-20, 0)

	if zero.LessThanZero() != false {
		t.Errorf("Expected %s to not be less than to 0, got %t", zero.String(), zero.LessThanZero())
	}

	if positive.LessThanZero() != false {
		t.Errorf("Expected %s to be < 0, got %t", positive.String(), positive.LessThanZero())
	}

	if negative.LessThanZero() != true {
		t.Errorf("Expected %s to not be less than 0, got %t", negative.String(), negative.LessThanZero())
	}
}
