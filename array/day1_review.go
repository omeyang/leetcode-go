package array

import (
	"math"
	"sort"
)

/*
 1. 两数之和  - 复习
*/

func twoSum2(nums []int, target int) []int {
	data := make(map[int]int)
	for index, value := range nums {
		if val, ok := data[target-value]; ok {
			return []int{val, index}
		}
		data[value] = index
	}
	return nil
}

/*
11. 盛水最多的容器 - 复习
*/
func maxArea2(height []int) int {
	ans, left, right := 0, 0, len(height)-1
	for left < right {
		area := (right - left) * compareMin(height[left], height[right])
		ans = compareMax(area, ans)
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return ans
}

/*
15 三数之和 - 复习
*/
func threeSum2(nums []int) (ans [][]int) {
	sort.Ints(nums)
	for index, value := range nums[:len(nums)-2] {
		left, right := index+1, len(nums)-1
		if value > 0 {
			break
		}
		if value+nums[index+1]+nums[index+2] > 0 {
			break
		}
		if nums[index]+nums[len(nums)-2]+nums[len(nums)-1] < 0 {
			continue
		}
		if index > 0 && value == nums[index-1] {
			continue
		}
		for left < right {
			if value+nums[left]+nums[right] < 0 {
				left++
			} else if value+nums[left]+nums[right] > 0 {
				right--
			} else {
				ans = append(ans, []int{value, nums[left], nums[right]})
				for left++; left < right && nums[left] == nums[left-1]; left++ {
				}
				for right--; right > left && nums[right] == nums[right+1]; right-- {
				}
			}
		}
	}
	return
}

/*
16.	最接近的三数之和 -复习
*/
func threeSumClosest2(nums []int, target int) (ans int) {
	sort.Ints(nums)
	ans = nums[0] + nums[1] + nums[len(nums)-1]
	for index, value := range nums[:len(nums)-2] {
		left, right := index+1, len(nums)-1
		for left < right {
			sum := value + nums[left] + nums[right]
			if math.Abs(float64(sum-target)) < math.Abs(float64(ans-target)) {
				ans = sum
			}
			if sum > target {
				right--
			} else {
				left++
			}
		}
	}
	return ans
}
