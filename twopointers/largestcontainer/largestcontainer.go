package largestcontainer

func largestcontainer(nums []int) int {
	maxArea := 0
	left, right := 0, len(nums)-1

	for left < right {
		height := min(nums[left], nums[right])
		length := right - left

		area := height * length
		if area > maxArea {
			maxArea = area
		}

		if nums[left] > nums[right] {
			right--
		} else if nums[left] < nums[right] {
			left++
		} else {
			right--
			left++
		}
	}
	return maxArea
}
