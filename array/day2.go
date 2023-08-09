package array

import (
	"math"
	"sort"
)

/*
16. 最接近的三数之和
*/
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[len(nums)-1]
	for index, value := range nums[:len(nums)-2] {
		left, right := index+1, len(nums)-1
		for left < right {
			sum := value + nums[left] + nums[right]
			if sum > target {
				right--
			} else {
				left++
			}
			if math.Abs(float64(sum-target)) < math.Abs(float64(res-target)) {
				res = sum
			}
		}
	}
	return res
}
