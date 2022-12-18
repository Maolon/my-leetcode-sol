package main

/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */

// @lc code=start
func searchInsert(nums []int, tgt int) int {
	l, h := 0, len(nums)-1

	for l <= h {
		
		mid := l + (h-l)>>1
		if nums[mid] >= tgt {
			h = mid - 1
		} else {
			// mid为最后一位
			// 或者因为是sort array
			// 所以只要mid + 1的值大于或者等于target 这就是插入的位置
			if (mid == len(nums)-1) || (nums[mid+1] >= tgt) {
				return mid + 1
			}
			l = mid + 1
		}
	}
	return 0
}

// @lc code=end
