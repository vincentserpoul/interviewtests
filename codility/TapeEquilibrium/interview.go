package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("toto\n")
}

func Solution(A []int) int {

	var lowestResults int

	//  brute force try all splits
	for i := 1; i < len(A); i++ {
		leftSum := 0
		rightSum := 0
		// sum all left side
		for j, mem := range A {
			if j < i {
				leftSum += mem
			}
			if j >= i {
				rightSum += mem
			}
		}
		diff := int(math.Abs(float64(leftSum - rightSum)))

		if i == 1 || diff < lowestResults {
			lowestResults = diff
		}
	}

	return lowestResults

}
