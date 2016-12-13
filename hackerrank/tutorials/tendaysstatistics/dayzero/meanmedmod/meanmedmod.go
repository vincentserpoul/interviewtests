package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {
	elementsCount, elements, err := getInputs()
	if err != nil {
		log.Fatal(err)
	}

	sort.Ints(elements)
	isMedianOneElement := (elementsCount%2 == 1)

	var mean, medianF float64
	var median int
	mode := -1
	previousElement := -1
	currentOccurences := 1
	var currentMaxOccurences int

	for i, element := range elements {
		// Mean
		mean += float64(element)

		// Median
		if isMedianOneElement {
			if i == (elementsCount / 2) {
				median = element
			}
		} else {
			if i == elementsCount/2-1 {
				median = element
			} else if i == (elementsCount / 2) {
				median += element
				medianF = float64(median) / float64(2)
			}
		}

		// Mode
		if previousElement == element {
			currentOccurences++
		} else {
			currentOccurences = 1

		}
		previousElement = element
		if currentOccurences > currentMaxOccurences {
			currentMaxOccurences = currentOccurences
			mode = element
		}

	}

	mean = float64(mean) / float64(elementsCount)
	fmt.Printf("%.1f\n%.1f\n%d\n", mean, medianF, mode)
}

func getInputs() (elementsCount int, elements []int, err error) {

	_, err = fmt.Scanf("%d", &elementsCount)
	if err != nil {
		return elementsCount, elements, err
	}

	for i := 0; i < elementsCount; i++ {
		var currentElement int
		_, err = fmt.Scanf("%d", &currentElement)
		if err != nil {
			return elementsCount, elements, err
		}
		elements = append(elements, currentElement)
	}

	if elementsCount == 0 || len(elements) == 0 {
		return elementsCount, elements, fmt.Errorf("no element given")
	}

	return elementsCount, elements, nil
}
