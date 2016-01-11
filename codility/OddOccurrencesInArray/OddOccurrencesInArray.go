package solution

import "sort"

// Solution returns the longest binary gap
func Solution(A []int) int {

	sort.IntSlice(A).Sort()
	previous := -1
	count := 0

	if len(A) == 0 {
		return A[0]
	}

	for _, elem := range A {
		if elem != previous && previous != -1 {
			if count%2 != 0 {
				return previous
			}
			count = 0
		}
		count++
		previous = elem
	}

	if count%2 != 0 {
		return A[len(A)-1]
	}

	return -1
}
