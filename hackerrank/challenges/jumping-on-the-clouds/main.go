package main

import (
	"fmt"
	"log"
)

func main() {
	clouds, err := getInputArray()
	if err != nil {
		log.Fatal(err)
	}
	minStepsCount, errCount := getMinStepCount(clouds)
	if errCount != nil {
		log.Fatal(errCount)
	}
	fmt.Printf("%d\n", minStepsCount)
}

func getMinStepCount(clouds []int) (int, error) {
	minStepCount := 0
	i := 0
	for i < len(clouds)-1 {
		minStepCount++
		if i+2 <= len(clouds)-1 && clouds[i+2] == 0 {
			i++
		}
		i++
	}

	return minStepCount, nil
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
