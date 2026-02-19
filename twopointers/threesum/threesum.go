package threesum

import (
	"slices"

	"github.com/cgeorgiades27/coding-patterns/twopointers/pairsum"
)

/*
Find all unique triplets in the array that sum to zero.

Args:
nums: List of integers

Returns:
List of triplets (as lists) that sum to zero, without duplicates

Time Complexity: O(n^2) - Sort O(n log n) + nested loop O(n^2)
Space Complexity: O(1) - excluding output array
*/
func threeSum(nums []int) [][]int {

	// sort first
	slices.Sort(nums)

	triplets := make([][]int, 0)

	for i := range nums {
		iVal := nums[i]
		result := pairsum.PairSum(nums[i+1:], -iVal)
		if len(result) < 2 {
			continue
		}
		triplet := append(result, iVal)
		slices.Sort(triplet)
		triplets = append(triplets, triplet)
	}

	return triplets
}
