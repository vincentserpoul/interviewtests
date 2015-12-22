package main

import "testing"

func Test(t *testing.T) {
	testCase := []int{2, 3, 1, 5}
	expectedSolution := 4

	if got := Solution(testCase); got != expectedSolution {
		t.Errorf("expecting %d, got %d", expectedSolution, got)
	}
}
