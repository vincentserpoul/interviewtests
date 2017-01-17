package main

import (
	"fmt"
	"log"
)

func main() {
	k, aList, err := getInput()
	if err != nil {
		log.Fatal(err)
	}

	divisibleSumPairCount := getDivisibleSumPairCount(k, aList)

	fmt.Printf("%d\n", divisibleSumPairCount)
}

func getDivisibleSumPairCount(k int, aList []int) (count int) {

	for i := 0; i < len(aList)-1; i++ {
		for j := i + 1; j <= len(aList)-1; j++ {
			if (aList[i]+aList[j])%k == 0 {
				count++
			}
		}
	}

	return count
}

func getInput() (int, []int, error) {
	var k, n int
	var err error
	_, err = fmt.Scanf("%d %d\n", &n, &k)
	if err != nil {
		return k, nil, err
	}

	aList := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &aList[i])
	}

	return k, aList, err
}
