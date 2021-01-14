package arraysandslices

// Sum adds up all numbers in the put array and returns the sum
func Sum(items []int) int {
	sum := 0

	for _, item := range items {
		sum += item
	}

	return sum
}
