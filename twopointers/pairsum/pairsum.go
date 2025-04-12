package pairsum

func PairSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{left, right}
		}
		if sum < target {
			left++
		}
		if sum > target {
			right--
		}
	}
	return []int{}
}
