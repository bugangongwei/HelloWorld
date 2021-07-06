package code

func SortedSquares(nums []int) []int {
	left := 0
	right := len(nums) - 1
	i := len(nums) - 1
	reuslt := make([]int, len(nums))

	for left < right {
		if nums[left]*nums[left] < nums[right]*nums[right] {
			reuslt[i] = nums[right] * nums[right]
			right--
		} else {
			reuslt[i] = nums[left] * nums[left]
			left++
		}

		i--
	}

	if left == right {
		reuslt[i] = nums[left] * nums[left]
	}
	return reuslt
}
