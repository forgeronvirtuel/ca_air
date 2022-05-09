package air

func QuicksortRecursive(arr []int, low, high int) []int {
	if low < high {
		pi := partition(arr, low, high, arr[high])
		Quicksort(arr, low, pi-1)
		Quicksort(arr, pi+1, high)

	}
	return arr

}

func Quicksort(arr []int, low, high int) []int {
	positions := make([][2]int, len(arr)*len(arr))
	positions[0][0] = low
	positions[0][1] = high
	for i, j, stop := 0, 0, false; !stop; i++ {
		low, high := positions[i][0], positions[i][1]
		if low == high {
			stop = true
			continue
		}
		pi := partition(arr, low, high, arr[high])
		j += 1
		if pi != low {
			positions[j][0] = low
			positions[j][1] = pi - 1
			j++
		}
		if pi != high {
			positions[j][0] = pi + 1
			positions[j][1] = high
		}
	}
	return arr
}

func partition(arr []int, low, high, pi int) int {
	j := low - 1
	for i := low; i <= high; i++ {
		if arr[i] <= pi {
			j++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	return j
}
