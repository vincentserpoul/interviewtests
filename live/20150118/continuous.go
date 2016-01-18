package solution

import "fmt"

// Solution returns the longest binary gap
func Solution(X []int, soughtSum int) bool {

	sum := 0
	potentialSum := []int{}

	for _, val := range X {
		sum += val
		potentialSum = append(potentialSum, val)
		fmt.Printf("appending %d, new sum == %d", val, sum)

		if sum > soughtSum {
			for sum > soughtSum {
				fmt.Printf(": sum too high, removing the first member, %d ", potentialSum[0])
				sum -= potentialSum[0]
				if len(potentialSum) == 1 {
					potentialSum = []int{}
				} else {
					potentialSum = potentialSum[1:len(potentialSum)]
				}
			}
		}

		if sum == soughtSum {
			fmt.Printf("all good, sum ==")
			return true
		}

		fmt.Printf(" >> sum: %d, new table: %v\n", sum, potentialSum)

	}

	return false
}
