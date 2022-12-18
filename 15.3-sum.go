package main

import "sort"

/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start
// map法
func threeSum(nums []int) [][]int {
	res := [][]int{}
	count := map[int]int{}
	sort.Ints(nums)
	// map去重
	for _, val := range nums {
		count[val]++
	}
	uniqNums := []int{}
	// 去重后的array
	for k := range count {
		uniqNums = append(uniqNums, k)
	}
	sort.Ints(uniqNums)

	for i := 0; i < len(uniqNums); i++ {
		// 判断三个值都为0的情况
		if uniqNums[i]*3 == 0 && count[uniqNums[i]] >= 3 {
			res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[i]})
		}

		for j := i + 1; j < len(uniqNums); j++ {
			// 2个i + 1个j = 0
			if (uniqNums[i]*2+uniqNums[j] == 0) && count[uniqNums[i]] > 1 {
				res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j]})
			}
			// 2个j + 1个i = 0
			if (uniqNums[i]+uniqNums[j] == 0) && count[uniqNums[j]] > 1 {
				res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[j]})
			}

			// 三个值不同，直接计算第三个值
			c := 0 - (uniqNums[i] + uniqNums[j])
			// 第三个值存在于数组
			if c > uniqNums[j] && count[c] > 0 {
				res = append(res, []int{uniqNums[i], uniqNums[j], c})
			}
		}
	}
	return res
}

// 双指针法
func threeSumPtr(nums []int) [][]int {
	sort.Ints(nums)
	res, start, end, index, addNum, length := make([][]int, 0), 0, 0, 0, 0, len(nums)
	// index = 中间值 start 为开始值 end为结束值
	for index = 1; index < length-1; index++ {
		start, end = 0, length-1
		if start < index && end > index {
			// 去除重复的数字
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < length-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			addNum = nums[start] + nums[index] + nums[end]
			if addNum == 0 {
				res = append(res, []int{nums[start], nums[index], nums[end]})
				start++
				end--
			} else if addNum > 0 {
				end--
			} else {
				start++
			}
		}

	}

	return res
}

// @lc code=end
