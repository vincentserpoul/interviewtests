package sorting

import "testing"

func Test_quickSort(t *testing.T) {
	qs := quickSort{}
	for _, tt := range getTestArrays() {
		t.Run(tt.name, func(t *testing.T) {
			qs.Sort(tt.arr)
			if !isSorted(tt.arr) {
				t.Errorf("quicksort: %v not sorted", tt.arr)
			}
		})
	}
}
