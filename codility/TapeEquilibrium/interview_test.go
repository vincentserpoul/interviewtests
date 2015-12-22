package main

import "testing"

func TestSolution(t *testing.T) {

	testArr := []int{3, 1, 2, 4, 3}

	if got := Solution(testArr); got != 1 {
		t.Errorf("expecting 1, got %d", got)
	}
}
