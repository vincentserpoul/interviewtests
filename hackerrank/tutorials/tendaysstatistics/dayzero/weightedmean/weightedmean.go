package main

import (
	"fmt"
	"log"
)

func main() {
	_, elements, weights, err := getInputs()
	if err != nil {
		log.Fatal(err)
	}

	var mean, divider float64
	for i, element := range elements {
		mean += float64(element) * float64(weights[i])
		divider += float64(weights[i])
	}
	fmt.Printf("%.1f", mean/divider)
}

func getInputs() (elementsCount int, elements []int, weights []int, err error) {

	_, err = fmt.Scanf("%d", &elementsCount)
	if err != nil {
		return elementsCount, elements, weights, err
	}

	for i := 0; i < elementsCount; i++ {
		var currentElement int
		_, err = fmt.Scanf("%d", &currentElement)
		if err != nil {
			return elementsCount, elements, weights, err
		}
		elements = append(elements, currentElement)
	}

	for i := 0; i < elementsCount; i++ {
		var currentweight int
		_, err = fmt.Scanf("%d", &currentweight)
		if err != nil {
			return elementsCount, elements, weights, err
		}
		weights = append(weights, currentweight)
	}

	if elementsCount == 0 || len(elements) == 0 || len(weights) == 0 {
		return elementsCount, elements, weights, fmt.Errorf("no element given")
	}

	return elementsCount, elements, weights, err
}
