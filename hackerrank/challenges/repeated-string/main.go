package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {

	var lilahS string
	var howManyLetters int
	_, err := fmt.Scanf("%s\n%d", &lilahS, &howManyLetters)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d\n", countHowManyAs(lilahS, howManyLetters))
}

func countHowManyAs(word string, howManyLetters int) int {
	leftLetters := howManyLetters % len(word)
	howManyFullRepeats := (howManyLetters - leftLetters) / len(word)

	return (howManyFullRepeats*strings.Count(word, "a") + strings.Count(word[0:leftLetters], "a"))
}
