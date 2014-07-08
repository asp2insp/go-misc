package algorithms

import "testing"
import "github.com/ca-geo/interview-prep/testutils"

func TestSort(t *testing.T) {
	input := []int{42, 5, 10, 36, 13}
	expected := []int{5, 10, 13, 36, 42}
	actual := RadixSort(input)
	testutils.CheckSlice(expected, actual, t)
}

func TestSort2(t *testing.T) {
	input := []int{52, 35, 0, 52, 64}
	expected := []int{0, 35, 52, 52, 64}
	actual := RadixSort(input)
	testutils.CheckSlice(expected, actual, t)
}
