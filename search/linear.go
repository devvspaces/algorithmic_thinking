package main

func LinearSearch(a []uint, value uint) bool {
	for _, v := range a {
		if v == value {
			return true
		}
	}
	return false
}

// func main() {

// 	searchArray := []uint{1, 2, 3, 4, 5, 6, 1, 5, 3, 10, 2, 5, 19}

// 	result := search(searchArray, searchValue)
// 	if result {
// 		fmt.Println("Found value in array")
// 	} else {
// 		fmt.Println("Value not found in array")
// 	}
// }
