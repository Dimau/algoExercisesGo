package binarysearch

import "math"

func FindElemIndex(elem int, array []int) (index int, success bool) {
	if len(array) == 0 {
		return 0, false
	}

	if len(array) == 1 {
		if elem == array[0] {
			return 0, true
		} else {
			return 0, false
		}
	}

	averageIndex := int(math.Floor(float64(len(array)-1) / 2))

	if array[averageIndex] == elem {
		return averageIndex, true
	}

	if array[averageIndex] < elem {
		return FindElemIndex(elem, array[averageIndex+1:])
	} else {
		return FindElemIndex(elem, array[:averageIndex])
	}
}
