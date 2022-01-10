package binarysearch

import "testing"

func TestHappyPath(t *testing.T) {
	array := []int{1, 2, 3, 4, 7, 10, 101, 140, 143, 178, 200}

	index, success := FindElemIndex(4, array)

	if index != 3 {
		t.Error("the right index for this case is 3")
	}

	if success != true {
		t.Error("in this case we should have success result")
	}
}

func TestElemIsNotInArray(t *testing.T) {
	array := []int{2, 3, 4, 7, 10}

	index, success := FindElemIndex(15, array)

	if index != 0 {
		t.Error("the right index for this case is 0 because the number is not in the array")
	}

	if success != false {
		t.Error("in this case we should have unsuccess result")
	}
}

func TestEmptyArray(t *testing.T) {
	array := []int{}

	index, success := FindElemIndex(4, array)

	if index != 0 {
		t.Error("index should be equal zero for empty array")
	}

	if success != false {
		t.Error("success should be fail for empty array")
	}
}

func TestArrayWithOneElementSuccess(t *testing.T) {
	array := []int{143}

	index, success := FindElemIndex(143, array)

	if index != 0 {
		t.Error("the right index is 0")
	}

	if success != true {
		t.Error("success should be true for this case")
	}
}

func TestArrayWithOneElementFail(t *testing.T) {
	array := []int{143}

	index, success := FindElemIndex(45, array)

	if index != 0 {
		t.Error("the right index is 0")
	}

	if success != false {
		t.Error("success should be false for this case")
	}
}
