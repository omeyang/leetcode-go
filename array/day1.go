package array

/*
1. 两数之和
*/
func twoSum(nums []int, target int) []int {
	data := make(map[int]int)
	for index, value := range nums {
		if val, ok := data[target-value]; ok {
			return []int{val, index}
		}
		data[value] = index
	}
	return nil
}
