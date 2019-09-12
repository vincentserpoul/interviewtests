package sorting

import "testing"

func Test_selectionSort(t *testing.T) {
	ss := selectionSort{}
	for _, tt := range getTestArrays() {
		t.Run(tt.name, func(t *testing.T) {
			ss.Sort(tt.arr)
			if !isSorted(tt.arr) {
				t.Errorf("selectionsort: %v not sorted", tt.arr)
			}
		})
	}
}
