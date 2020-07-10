package _3_rotatedarray

func search(nums []int, target int) int {
	if len(nums) <= 0 {
		return -1
	}

	left := 0
	right := len(nums)-1
	mid := 0

	for left <= right {
		mid = (left+right) / 2
		midNum := nums[mid]
		if midNum == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			if nums[0] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[len(nums) - 1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
}