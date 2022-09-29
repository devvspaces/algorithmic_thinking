package utils

import (
	"fmt"
	"testing"
)

func TestEqual(t *testing.T) {

	tests := [][]int{
		{1, 2, 3, 4, 5, 6},
		{1, 2, 3, 4, 5, 7},
		{1, 2, 3, 4, 5, 7, 6},
	}

	compares := [][]int{
		{1, 2, 3, 4, 5, 6},
		{1, 2, 3, 4, 5, 6},
		{1, 2, 3, 4, 5, 6},
	}

	expects := []bool{
		true,
		false,
		false,
	}

	for idx, value := range tests {
		name := fmt.Sprintf("EqualTest_%v", idx)
		t.Run(name, func(sub *testing.T) {
			compare := compares[idx]
			expected := expects[idx]
			if ok := Equal(value, compare); ok != expected {
				t.Errorf("Expected %v for %v == %v got %v", expected, value, compare, ok)
			}
		})
	}

}
