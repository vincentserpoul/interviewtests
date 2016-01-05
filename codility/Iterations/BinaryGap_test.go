package solution

import "testing"

func TestSolution(t *testing.T) {
	c := []struct {
		i         int
		expectedO int
	}{
		{i: 9, expectedO: 2},
		{i: 1041, expectedO: 5},
		{i: 529, expectedO: 4},
		{i: 20, expectedO: 1},
		{i: 15, expectedO: 0},
		{i: 1, expectedO: 0},
	}

	for _, testedCase := range c {
		gotO := Solution(testedCase.i)
		if gotO != testedCase.expectedO {
			t.Errorf("Solution(%d) expected %d, got %d", testedCase.i, testedCase.expectedO, gotO)
			return
		}
	}

}
