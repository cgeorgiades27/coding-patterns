package pairsum

import (
	"slices"
	"testing"
)

var pairSumTests = []struct {
	target   int
	nums     []int
	expected []int
}{
	{
		7,
		[]int{-5, -2, 3, 4, 6},
		[]int{2, 3},
	},
	{
		2,
		[]int{1, 1, 1},
		[]int{0, 2},
	},
	{
		0,
		[]int{},
		[]int{},
	},
	{
		1,
		[]int{1},
		[]int{},
	},
	{
		5,
		[]int{2, 3},
		[]int{0, 1},
	},
	{
		5,
		[]int{2, 2, 3},
		[]int{0, 2},
	},
	{
		2,
		[]int{-1, 2, 3},
		[]int{0, 2},
	},
	{
		-5,
		[]int{-3, -2, -1},
		[]int{0, 1},
	},
}

func TestPairSum(t *testing.T) {
	for index, input := range pairSumTests {
		actual := PairSum(input.nums, input.target)
		result := slices.Equal(input.expected, actual)
		if !result {
			t.Errorf("test %d failed, expected: %v, got %v", index, input.expected, actual)
		}
	}
}
