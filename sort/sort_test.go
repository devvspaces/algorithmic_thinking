package sort

import (
	"algo/utils"
	"fmt"
	"testing"
)

func HelpTestSort(t *testing.T, sort func([]int)) {
	arrs := [][]int{
		{1, 5, 4, 6, 2, 7, 5, 4},
		{19, 4, 2, 8, 12, 11, 20, 7},
		{7, 8, 9, 4, 2, 6, 5, 3},
	}
	res := [][]int{
		{1, 2, 4, 4, 5, 5, 6, 7},
		{2, 4, 7, 8, 11, 12, 19, 20},
		{2, 3, 4, 5, 6, 7, 8, 9},
	}
	for idx := range arrs {
		question := make([]int, len(arrs[idx]))
		copy(question, arrs[idx])
		expected := res[idx]
		test_name := fmt.Sprintf("test_%v", idx)
		t.Run(test_name, func(sub *testing.T) {
			sort(question)
			if utils.Equal(question, expected) == false {
				sub.Errorf("Sorting array %v gives us %v, expected %v", arrs[idx], question, expected)
			}
		})
	}
}

func TestQuicksortSort(t *testing.T) {
	HelpTestSort(t, Quicksort)
}

func TestBubbleSortSort(t *testing.T) {
	HelpTestSort(t, BubbleSort)
}
