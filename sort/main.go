package main

import "fmt"

func bubble_sort(a []int) {
	for i := 0; i < len(a); i++ {
		for index := 0; index < len(a)-1-i; index++ {
			index2 := index + 1
			if a[index] > a[index2] {
				val := a[index2]
				a[index2] = a[index]
				a[index] = val
			}
		}
	}
}

func main() {
	fmt.Println("This is a bubble sort algorithm")
	arr := []int{1, 3, 2, 5, 3, 7, 4, 9, 8, 6}
	bubble_sort(arr)
	fmt.Println(arr)

}
