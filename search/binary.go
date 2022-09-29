package main

import (
	"flag"
	"fmt"
	"sort"
)

func search(a []uint, v uint) bool {

	high := len(a)
	low := 0

	for low < high {
		index := ((high - low) / 2) + low
		ivalue := a[index]

		if ivalue == v {
			return true
		}

		if ivalue > v {
			high = index
			continue
		}

		low = index + 1
	}

	return false
}

var searchValue uint

func init() {
	flag.UintVar(&searchValue, "s", 0, "search for the number in array")
}

type SortBy []uint

func (a SortBy) Len() int {
	return len(a)
}
func (a SortBy) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a SortBy) Less(i, j int) bool {
	return a[i] < a[j]
}

func main() {

	flag.Parse()

	searchArray := SortBy{1, 2, 3, 4, 5, 6, 5, 3, 10, 2, 5, 19}
	sort.Sort(searchArray)

	fmt.Println(searchArray)

	result := search(searchArray, searchValue)
	if result {
		fmt.Println("Found value in array")
	} else {
		fmt.Println("Value not found in array")
	}
}
