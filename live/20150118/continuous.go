package solution

// Solution returns the longest binary gap
func Solution(X []int, soughtSum int) bool {

	sum := 0
	lower := 0
	upper := 0

	for upper <= len(X) && lower < len(X) {

		if sum < soughtSum {
			sum += X[upper]
			upper++
		} else if sum > soughtSum {
			sum -= X[lower]
			lower++
		}
		if sum == soughtSum {
			return true
		}
	}

	return false
}
