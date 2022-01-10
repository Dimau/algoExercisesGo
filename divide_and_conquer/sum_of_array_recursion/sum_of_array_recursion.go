package sum_of_array_recursion

func SumArray(array []int) int {
	if len(array) == 0 {
		return 0
	}
	if len(array) == 1 {
		return array[0]
	} else {
		return array[0] + sumArray(array[1:])
	}
}
