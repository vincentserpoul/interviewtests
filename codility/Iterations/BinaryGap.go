package solution

import "fmt"

// Solution returns the longest binary gap
func Solution(N int) int {
	str := fmt.Sprintf("%b", N)
	var count int
	var maxCount int

	for i := 0; i < len(str); i++ {
		if i > 0 && string(str[i-1]) == "1" && string(str[i]) == "0" {
			count++
			continue
		}
		if count > 0 {
			if string(str[i]) == "0" {
				count++
			} else {
				if count > maxCount {
					maxCount = count
					count = 0
				}
			}

		}
	}
	return maxCount
}
