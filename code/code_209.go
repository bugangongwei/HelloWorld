package code

// 2, 3, 1, 2, 4, 3
func MinSubArrayLen(target int, nums []int) int {
	i := 0
	sum := 0
	result := len(nums) + 1

	for j := 0; j < len(nums); j++ {
		sum += nums[j]

		for sum >= target {
			wLen := j - i + 1
			if wLen < result {
				result = wLen
			}
			sum -= nums[i]
			i++
		}
	}

	if result == len(nums)+1 {
		return 0
	} else {
		return result
	}
}
