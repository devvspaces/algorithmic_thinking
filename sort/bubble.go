package sort

func BubbleSort(a []int) {
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
