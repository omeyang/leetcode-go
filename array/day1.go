package array

import "sort"

/*
1. 两数之和
时间复杂度  O(n)
空间复杂度  O(n)
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

/*
11. 盛水最多的容器
时间复杂度  O(n)
空间复杂度  O(1)
*/
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	ans := 0
	for left < right {
		area := (right - left) * compareMin(height[left], height[right])
		ans = compareMax(ans, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return ans
}

func compareMax(curr, data int) int {
	if curr > data {
		return curr
	}
	return data
}

func compareMin(curr, data int) int {
	if curr > data {
		return data
	}
	return curr
}

/*
15. 三数之和
时间复杂度  O(n+ nlog(n))
空间复杂度  O(1)
*/
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := [][]int{}
	for index, value := range nums[:len(nums)-2] {
		if value > 0 {
			break
		}
		if value+nums[index+1]+nums[index+2] > 0 {
			break
		}
		if value+nums[len(nums)-2]+nums[len(nums)-1] < 0 {
			continue
		}
		if index > 0 && value == nums[index-1] { // 跳过重复数字
			continue
		}

		left, right := index+1, len(nums)-1
		for left < right {
			if value+nums[left]+nums[right] > 0 {
				right--
			} else if value+nums[left]+nums[right] < 0 {
				left++
			} else {
				ans = append(ans, []int{value, nums[left], nums[right]})
				for left++; left < right && nums[left] == nums[left-1]; left++ {
				}
				for right--; right > left && nums[right] == nums[right+1]; right-- {
				}
			}
		}
	}
	return ans
}