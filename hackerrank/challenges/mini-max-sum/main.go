package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {

	inputs := make([]int, 5)
	_, err := fmt.Scanf("%d %d %d %d %d", &inputs[0], &inputs[1], &inputs[2], &inputs[3], &inputs[4])
	if err != nil {
		log.Fatal(err)
	}

	sort.Ints(inputs)

	fmt.Printf("%d %d\n",
		inputs[0]+inputs[1]+inputs[2]+inputs[3],
		inputs[1]+inputs[2]+inputs[3]+inputs[4],
	)
}
