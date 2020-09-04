package _5_3sum

import "sort"

func threeSUm(nums []int) [][]int {
	result := make([][]int, 0)
	length := len(nums)
	if length < 3 {
		return result
	}
	sort.Ints(nums)
	for i := 0; i < length; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		k := length - 1
		for j := i + 1; j < length; j++ {
			if j > i + 1 && nums[j] == nums[j-1] {
				continue
			}
			for ; j < k && nums[j] + nums[k] > -1 * nums[i]; k-- {
			}
			if j == k {
				break
			}
			if nums[j]+nums[k] == -1*nums[i] {
				result = append(result, []int{nums[i], nums[j], nums[k]})
			}
		}
	}
	return result
}