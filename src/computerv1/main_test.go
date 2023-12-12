package main

import "testing"

func TestGivenInput(t *testing.T) {
	result := givenInput("5 * X^0 + 4 * X^1 - 9.3 * X^2 = 1 * X^0")
	expected := " 4 * X^0 + 4 * X^1 - 9.3 * X^2 = 0"

	if(result != expected) {
		t.Errorf("returned %s, expected %s", result, expected)
	}
}