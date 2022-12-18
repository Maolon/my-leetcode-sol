package main

/*
 * @lc app=leetcode id=189 lang=golang
 *
 * [189] Rotate Array
 */

// @lc code=start
// O(1) space solution
func rotate(nums []int, k int) {
	k %= len(nums)
	//[1,2,3,4,5,6,7]
	reverse(nums)
	//[7 6 5 4 3 2 1]
	reverse(nums[:k])
	//[5 6 7 4 3 2 1]
	reverse(nums[k:])
	//[5,6,7,1,2,3,4]
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

// O(n) space 解法
func rotateSol2(nums []int, k int) {
	l := len(nums)
	bk := make([]int, l)
	//copy一个nums
	copy(bk, nums)

	for i, n := range bk {
		j := (i + k) % l
		nums[j] = n
	}
}

// @lc code=end
