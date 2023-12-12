package main

import "testing"

func TestGivenInputA(t *testing.T) {
	reduced, degree := givenInput("5 * X^0 + 4 * X^1 - 9.3 * X^2 = 1 * X^0")
	expectedReduce := "4 * X^0 + 4 * X^1 - 9.3 * X^2 = 0"
	expectedDegree := 2

	if(reduced != expectedReduce) {
		t.Errorf("returned: %s\n expected: %s", expectedReduce, expectedReduce)
	}
	if(degree != expectedDegree) {
		t.Errorf("returned: %d\n expected: %d", degree, expectedDegree)
	}
}

func TestGivenInputB(t *testing.T) {
	reduced, degree := givenInput("5 * X^0 + 4 * X^1 = 4 * X^0")
	expectedReduce := "1 * X^0 + 4 * X^1 = 0"
	expectedDegree := 1

	if(reduced != expectedReduce) {
		t.Errorf("returned: %s\n expected: %s", expectedReduce, expectedReduce)
	}
	if(degree != expectedDegree) {
		t.Errorf("returned: %d\n expected: %d", degree, expectedDegree)
	}
}

func TestGivenInputC(t *testing.T) {
	reduced, degree := givenInput("8 * X^0 - 6 * X^1 + 0 * X^2 - 5.6 * X^3 = 3 * X^0")
	expectedReduce := "5 * X^0 - 6 * X^1 + 0 * X^2 - 5.6 * X^3 = 0"
	expectedDegree := 3

	if(reduced != expectedReduce) {
		t.Errorf("returned: %s\n expected: %s", expectedReduce, expectedReduce)
	}
	if(degree != expectedDegree) {
		t.Errorf("returned: %d\n expected: %d", degree, expectedDegree)
	}
}