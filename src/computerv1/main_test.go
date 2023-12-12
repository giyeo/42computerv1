package main

import "testing"

func TestGivenInput(t *testing.T) {
	testCases := []struct {
		input           string
		expectedReduce  string
		expectedDegree  int
		expectedSolution string
	}{
		{
			input:           "5 * X^0 + 4 * X^1 - 9.3 * X^2 = 1 * X^0",
			expectedReduce:  "4 * X^0 + 4 * X^1 - 9.3 * X^2 = 0",
			expectedDegree:  2,
			expectedSolution: "",
		},
		{
			input:           "5 * X^0 + 4 * X^1 = 4 * X^0",
			expectedReduce:  "1 * X^0 + 4 * X^1 = 0",
			expectedDegree:  1,
			expectedSolution: "The solution is:\n-0.25",
		},
		{
			input:           "8 * X^0 - 6 * X^1 + 0 * X^2 - 5.6 * X^3 = 3 * X^0",
			expectedReduce:  "5 * X^0 - 6 * X^1 + 0 * X^2 - 5.6 * X^3 = 0",
			expectedDegree:  3,
			expectedSolution: "The polynomial degree is strictly greater than 2, I can't solve.",
		},
		{
			input:           "42 * X^0 = 42 * X^0",
			expectedReduce:  "",
			expectedDegree:  0,
			expectedSolution: "True",
		},
		{
			input:           "42 * X^0 = 43 * X^0",
			expectedReduce:  "",
			expectedDegree:  0,
			expectedSolution: "False",
		},
	}

	for _, tc := range testCases {
		reduced, degree, solution := givenInput(tc.input)

		println(tc.expectedReduce + "\n" + reduced)
		t.Run(tc.input, func(t *testing.T) {
			t.Helper()

			if tc.expectedReduce != "" && reduced != tc.expectedReduce {
				t.Errorf("Reduced: %s\nExpected: %s", reduced, tc.expectedReduce)
			}
			if tc.expectedDegree != -1 && degree != tc.expectedDegree {
				t.Errorf("Degree: %d\nExpected: %d", degree, tc.expectedDegree)
			}
			if tc.expectedSolution != "" && solution != tc.expectedSolution {
				t.Errorf("Solution: %s\nExpected: %s", solution, tc.expectedSolution)
			}
		})
	}
}
