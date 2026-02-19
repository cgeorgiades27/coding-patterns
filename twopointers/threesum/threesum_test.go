package threesum

import (
	"reflect"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name:     "basic case with one triplet",
			nums:     []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name:     "empty array",
			nums:     []int{},
			expected: [][]int{},
		},
		{
			name:     "array too small",
			nums:     []int{0, 1},
			expected: [][]int{},
		},
		{
			name:     "no triplets sum to zero",
			nums:     []int{1, 2, 3, 4, 5},
			expected: [][]int{},
		},
		{
			name:     "all zeros",
			nums:     []int{0, 0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
		{
			name:     "multiple triplets",
			nums:     []int{-2, 0, 1, 1, 2},
			expected: [][]int{{-2, 0, 2}, {-2, 1, 1}},
		},
		{
			name:     "duplicates in input",
			nums:     []int{-1, 0, 1, 0, -1, 1},
			expected: [][]int{{-1, 0, 1}},
		},
		{
			name:     "larger array",
			nums:     []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6},
			expected: [][]int{{-4, -2, 6}, {-4, 0, 4}, {-4, 1, 3}, {-4, 2, 2}, {-2, -2, 4}, {-2, 0, 2}},
		},
		{
			name:     "single zero",
			nums:     []int{0},
			expected: [][]int{},
		},
		{
			name:     "exactly three elements summing to zero",
			nums:     []int{-1, 0, 1},
			expected: [][]int{{-1, 0, 1}},
		},
		{
			name:     "exactly three elements not summing to zero",
			nums:     []int{1, 2, 3},
			expected: [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := threeSum(tt.nums)

			// Sort both result and expected for comparison
			sortTriplets(result)
			sortTriplets(tt.expected)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("threeSum(%v) = %v, want %v", tt.nums, result, tt.expected)
			}
		})
	}
}

// Helper function to sort triplets for comparison
func sortTriplets(triplets [][]int) {
	// Sort each individual triplet
	for _, triplet := range triplets {
		sort.Ints(triplet)
	}

	// Sort the slice of triplets
	sort.Slice(triplets, func(i, j int) bool {
		for k := 0; k < 3 && k < len(triplets[i]) && k < len(triplets[j]); k++ {
			if triplets[i][k] != triplets[j][k] {
				return triplets[i][k] < triplets[j][k]
			}
		}
		return false
	})
}
