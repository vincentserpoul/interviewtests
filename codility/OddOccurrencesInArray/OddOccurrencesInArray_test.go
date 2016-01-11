package solution

import "testing"

func TestSolution(t *testing.T) {
	c := []struct {
		A              []int
		expectedOutput int
	}{
		{A: []int{9, 3, 9, 3, 9, 7, 9}, expectedOutput: 7},
		{A: []int{42}, expectedOutput: 42},
		{A: []int{42, 42, 42}, expectedOutput: 42},
		{A: []int{42, 42, 42, 42}, expectedOutput: -1},
	}

	for _, testedCase := range c {
		gotO := Solution(testedCase.A)
		if gotO != testedCase.expectedOutput {
			t.Errorf("Solution(%v) expected %d, got %d", testedCase.A, testedCase.expectedOutput, gotO)
			return
		}
	}

}
