package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {
	socks, err := getSocks()
	if err != nil {
		log.Fatal(err)
	}

	pairCount := pairSocks(socks)

	fmt.Printf("%d\n", pairCount)
}

func getSocks() ([]int, error) {
	var socksCount int
	var err error
	_, err = fmt.Scanf("%d\n", &socksCount)
	if err != nil {
		return nil, err
	}

	socks := make([]int, socksCount)

	for i := 0; i < socksCount; i++ {
		fmt.Scanf("%d", &socks[i])
	}

	return socks, err
}

func pairSocks(socks []int) (pairCount int) {
	sort.Ints(socks)
	previousSockColor := socks[0]

	for i := 1; i < len(socks); i++ {
		if socks[i] == previousSockColor {
			pairCount++
			previousSockColor = 0
		} else {
			previousSockColor = socks[i]
		}
	}

	return pairCount
}
