package code

// nums = [0,1,2,2,3,0,4,2], val = 2  5
// nums = [3,2,2,3], val = 3 2
func RemoveElement(nums []int, val int) int {
	left := 0
	for _, v := range nums {
		if v != val {
			nums[left] = v
			left++
		}
	}
	return left
}
