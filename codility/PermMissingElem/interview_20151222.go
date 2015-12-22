package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("1")
}

func Solution(a []int) int {
	sort.Sort(intSlice(a))
	for i, val := range a {
		if i+1 != val {
			return i + 1
		}
	}

	return len(a)
}

type intSlice []int

func (a intSlice) Len() int           { return len(a) }
func (a intSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a intSlice) Less(i, j int) bool { return a[i] < a[j] }
