package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	var nDays int

	_, err := fmt.Scanf("%d", &nDays)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(countLikedByPeople(nDays))
}

func countLikedByPeople(nDays int) (totalLikedCount int) {

	sentToCount := 5
	for i := 1; i <= nDays; i++ {
		likedByCount := int(math.Floor(float64(sentToCount) / float64(2)))
		totalLikedCount += likedByCount

		// log.Printf(
		// 	"day %d: sent to %d, liked by %d (adding %d more)",
		// 	i,
		// 	sentToCount,
		// 	totalLikedCount,
		// 	likedByCount,
		// )
		sentToCount = likedByCount * 3
	}
	return totalLikedCount
}
