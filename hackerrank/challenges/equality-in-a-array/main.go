package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {
	a, err := getInputArray()
	if err != nil {
		log.Fatal(err)
	}

	minDel := getMinimumElemDelete(a)

	fmt.Printf("%d\n", minDel)
}

func getMinimumElemDelete(a []int) int {
	sort.Ints(a)
	maxOccurence := 1
	occurence := 1
	var previousAi int

	for i, ai := range a {
		if i > 0 {
			if ai == previousAi {
				occurence++
			} else {
				occurence = 1
			}
			if occurence > maxOccurence {
				maxOccurence = occurence
			}
		}

		previousAi = ai
	}

	return len(a) - maxOccurence
}

func getInputArray() (a []int, err error) {
	var n int
	_, err = fmt.Scanf("%d\n", &n)
	if err != nil {
		return a, err
	}

	for i := 0; i < n; i++ {
		var ai int
		_, erri := fmt.Scanf("%d", &ai)
		if erri != nil {
			return a, erri
		}
		a = append(a, ai)
	}

	return a, nil
}
