package main

import "testing"

func TestGetNumberOfPayments(t *testing.T) {
	expected := 79

	result, _, e := getNumberOfPayments(10000.00, 0.05, 150.00, 1)

	if e != nil {
		t.Errorf("%v", e)
	}

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
