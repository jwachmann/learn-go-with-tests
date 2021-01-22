package arraysandslices

// Sum adds up all numbers in the put array and returns the sum
func Sum(items []int) int {
	sum := 0

	for _, item := range items {
		sum += item
	}

	return sum
}

// SumAll returns an array where each index is the sum of the items at that index in the input arrays
func SumAll(listOfListsOfNumbers ...[]int) []int {
	var sums []int

	for _, numbers := range listOfListsOfNumbers {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

// SumAllTails sums up the numbers other than the first item in each input array
func SumAllTails(listOfListsOfNumbers ...[]int) []int {
	var sums []int

	for _, numbers := range listOfListsOfNumbers {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
