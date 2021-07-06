package code

func Search(nums []int, target int) int {
	// [left, right]
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		if target == nums[mid] {
			return mid
		}

		// [mid, right]
		if target > nums[mid] {
			left = mid + 1
		}

		// [left, mid]
		if target < nums[mid] {
			right = mid - 1
		}
	}

	return -1
}

/*
升级版
如果没找到 target, 就返回 target 按顺序插入的顺序
*/
func SearchOrInsert(nums []int, target int) int {
	// [left, right]
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := (left + right) / 2

		if target == nums[mid] {
			return mid
		}

		// [mid, right]
		if target > nums[mid] {
			left = mid + 1
		}

		// [left, mid]
		if target < nums[mid] {
			right = mid - 1
		}
	}

	if left == right {
		if target <= nums[left] {
			return left
		}

		if target > nums[right] {
			return right + 1
		}
	}

	return -1
}
