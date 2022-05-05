package air

import "sort"

func IsAllDuplicated(numbers []int) (bool, int) {
	sort.Ints(numbers)
	for i := 0; i < len(numbers); i += 2 {
		if numbers[i] != numbers[i+1] {
			return false, numbers[i]
		}
	}
	return true, 0
}
