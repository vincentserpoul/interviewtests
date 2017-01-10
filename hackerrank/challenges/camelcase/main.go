package main

import (
	"fmt"
	"log"
	"unicode"
)

func main() {
	wordCount, err := howManyWords()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d\n", wordCount)
}

func howManyWords() (wordCount int, err error) {
	var sentence string
	_, err = fmt.Scanf("%s", &sentence)
	if err != nil {
		return wordCount, err
	}

	wordCount = countWords(sentence)
	return wordCount, nil
}

func countWords(s string) int {
	if len(s) == 0 {
		return 0
	}

	wordCount := 1
	for _, runeValue := range s {
		if unicode.IsUpper(runeValue) {
			wordCount++
		}
	}
	return wordCount
}
