package sort

func partition(arr []int, lo int, hi int) int {
	pivot := arr[hi]
	idx := lo - 1

	for i := lo; i < hi; i++ {
		if arr[i] <= pivot {
			idx++
			temp := arr[i]
			arr[i] = arr[idx]
			arr[idx] = temp
		}
	}

	idx++
	arr[hi] = arr[idx]
	arr[idx] = pivot

	return idx

}

func qs(arr []int, lo int, hi int) {

	if lo >= hi {
		return
	}

	idx := partition(arr, lo, hi)

	qs(arr, lo, idx-1)
	qs(arr, idx+1, hi)

}

func Quicksort(arr []int) {
	low := 0
	high := len(arr) - 1
	qs(arr, low, high)
}
