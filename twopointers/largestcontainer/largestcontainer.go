package largestcontainer

func largestcontainer(nums []int) int {
	left, right := 0, len(nums)-1
	maxArea := 0
	for left < right {
		heightLeft, heightRight := nums[left], nums[right]
		area := min(heightLeft, heightRight) * (right - left)
		maxArea = max(maxArea, area)

		// move the pointers
		if heightLeft > heightRight {
			right--
		} else if heightLeft < heightRight {
			left++
		} else {
			left++
			right--
		}
	}
	return maxArea
}
