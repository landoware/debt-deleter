package debts_test

import (
	"testing"

	"github.com/landoware/debt-deleter/debts"
	"github.com/landoware/debt-deleter/interest"
	"github.com/landoware/debt-deleter/money"
)

func TestPayOnLoanWithoutInterest(t *testing.T) {
	loan := debts.Loan{
		Name:           "test",
		Principal:      money.NewMoney(100, 0),
		UnpaidInterest: money.NewMoney(0, 0),
		Rate:           interest.NewRateFromParts(5, 0),
		MinPayment:     money.NewMoney(50, 0),
		DueDay:         1,
	}

	remainder := loan.PayOnLoan(loan.MinPayment)
	expected := money.NewMoney(50, 0)

	if remainder.NotEqualsZero() {
		t.Errorf("Expected no remainder, got %s", remainder.String())
	}

	if loan.Principal.NotEquals(expected) {
		t.Errorf("Expected balance after %s payment to be %s, %s returned", loan.MinPayment, expected, loan.Principal.String())
	}
}

func TestPayOnLoanWithInterest(t *testing.T) {
	loan := debts.Loan{
		Name:           "test",
		Principal:      money.NewMoney(100, 0),
		UnpaidInterest: money.NewMoney(25, 0),
		Rate:           interest.NewRateFromParts(5, 0),
		MinPayment:     money.NewMoney(50, 0),
		DueDay:         1,
	}

	remainder := loan.PayOnLoan(loan.MinPayment)
	expected := money.NewMoney(75, 0)

	if remainder.NotEqualsZero() {
		t.Errorf("Expected no remainder, got %s", remainder.String())
	}

	if loan.Principal.NotEquals(expected) {
		t.Errorf("Expected balance after %s payment to be %s, %s returned", loan.MinPayment, expected, loan.Principal)
	}
}

func TestOverpayOnLoanWithoutInterest(t *testing.T) {
	loan := debts.Loan{
		Name:           "test",
		Principal:      money.NewMoney(100, 0),
		UnpaidInterest: money.NewMoney(0, 0),
		Rate:           interest.NewRateFromParts(5, 0),
		MinPayment:     money.NewMoney(150, 0),
		DueDay:         1,
	}

	remainder := loan.PayOnLoan(loan.MinPayment)
	expected := money.NewMoney(50, 0)

	if remainder.NotEquals(expected) {
		t.Errorf("Expected remainder of %s, got %s", expected.String(), remainder.String())
	}

	if loan.Principal.NotEqualsZero() {
		t.Errorf("Expected balance after overpayment to be $0.00, %s returned", loan.Principal.String())
	}
}

func TestOverpayOnLoanWithInterest(t *testing.T) {
	loan := debts.Loan{
		Name:           "test",
		Principal:      money.NewMoney(100, 0),
		UnpaidInterest: money.NewMoney(25, 0),
		Rate:           interest.NewRateFromParts(5, 0),
		MinPayment:     money.NewMoney(150, 0),
		DueDay:         1,
	}

	remainder := loan.PayOnLoan(loan.MinPayment)
	expected := money.NewMoney(25, 0)

	if remainder.NotEquals(expected) {
		t.Errorf("Expected remainder of %s, got %s", expected.String(), remainder.String())
	}

	if loan.Principal.NotEqualsZero() {
		t.Errorf("Expected balance after overpayment to be $0.00, %s returned", loan.Principal.String())
	}
}

func TestPaymentLessThanInterest(t *testing.T) {
	loan := debts.Loan{
		Name:           "test",
		Principal:      money.NewMoney(100, 0),
		UnpaidInterest: money.NewMoney(50, 0),
		Rate:           interest.NewRateFromParts(5, 0),
		MinPayment:     money.NewMoney(25, 0),
		DueDay:         1,
	}

	remainder := loan.PayOnLoan(loan.MinPayment)
	expected := money.NewMoney(125, 0)

	if remainder.NotEqualsZero() {
		t.Errorf("Expected no remainder, got %s", remainder.String())
	}

	if loan.Principal.Add(loan.UnpaidInterest).NotEquals(expected) {
		t.Errorf("Expected total balance after %s payment to be %s, %s returned", loan.MinPayment, expected, loan.Principal.Add(loan.UnpaidInterest))
	}

}
