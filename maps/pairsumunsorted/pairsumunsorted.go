package pairsumunsorted

func pairsumunsorted(nums []int, target int) []int {
	complimentMap := make(map[int]int)
	for i, num := range nums {
		compliment := target - num
		if complimentIndex, ok := complimentMap[compliment]; ok {
			return []int{complimentIndex, i}
		}
		complimentMap[num] = i
	}
	return []int{}
}
