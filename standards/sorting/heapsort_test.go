package sorting

import "testing"

func Test_heapSort(t *testing.T) {
	qs := heapSort{}
	for _, tt := range getTestArrays() {
		t.Run(tt.name, func(t *testing.T) {
			qs.Sort(tt.arr)
			if !isSorted(tt.arr) {
				t.Errorf("heapsort: %v not sorted", tt.arr)
			}
		})
	}
}
