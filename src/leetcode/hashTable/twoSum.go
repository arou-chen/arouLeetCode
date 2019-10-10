/**
 *Create by chensr on 2019/9/24
*/

package hashTable

func TwoSum(nums []int, target int) []int {
	tempMap := make(map[int]int)
	var a, b int
	for idx, data := range nums {
		tempMap[data] = idx
	}
	for idx, data := range nums {
		key := target - data
		num, ok := tempMap[key]
		if ok && num != idx {
			a = num
			b = idx
			break
		}
	}
	return []int{a, b}
}