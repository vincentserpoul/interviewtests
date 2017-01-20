package main

import (
	"fmt"
	"log"
)

func main() {
	n, k, items, amountAsked, err := getInput()
	if err != nil {
		log.Fatal(err)
	}

	amount := getHowmuchBrianOwes(n, k, items, amountAsked)

	if amount == 0 {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Printf("%d\n", amount)
	}

}

func getInput() (int, int, []int, int, error) {
	var n, k int
	var amountAsked int
	var err error

	// get n and k
	_, err = fmt.Scanf("%d %d\n", &n, &k)
	if err != nil {
		return n, k, nil, amountAsked, err
	}

	// get item list
	itemList := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &itemList[i])
	}

	// get asked amount
	_, err = fmt.Scanf("%d\n", &amountAsked)
	if err != nil {
		return n, k, nil, amountAsked, err
	}

	return n, k, itemList, amountAsked, err
}

func getHowmuchBrianOwes(n int, k int, items []int, amountAsked int) (amount int) {

	var sumForAnna int

	for i, itemPrice := range items {
		if i == k {
			continue
		}
		sumForAnna += itemPrice
	}

	return amountAsked - sumForAnna/2
}
