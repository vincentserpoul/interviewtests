package solution

import "testing"

func TestSolution(t *testing.T) {

	c := []struct {
		intArray       []int
		soughtSum      int
		expectedResult bool
	}{
		{[]int{2000, 5, 4, 7, 2, 12}, 21, true},
		{[]int{1, 4, 4, 23, 2}, 8, true},
		{[]int{2, 56, 5, 3, 2}, 1, false},
	}

	for _, testc := range c {
		if got := Solution(testc.intArray, testc.soughtSum); got != testc.expectedResult {
			t.Errorf("Solution(%v, %d) expecting %t, got %t", testc.intArray, testc.soughtSum, testc.expectedResult, got)
		}
	}
}
