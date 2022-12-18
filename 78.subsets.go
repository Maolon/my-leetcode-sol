package main

/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */

// @lc code=start
func subsets(nums []int) [][]int {
	res := [][]int{}
	dfs(nums, []int{}, &res)
	return res
}

func dfs(nums, path []int, res *[][]int) {
	b := make([]int, len(path))
	copy(b, path)
	*res = append(*res, b)
	for i := 0; i < len(nums); i++ {
		// last n elements of slice
		dfs(nums[i+1:], append(b, nums[i]), res)
	}
}

// 直接生成subset
func subsetsSol2(nums []int) [][]int {
	res := make([][]int, 0)
	res = append(res, []int{})

	for _, n := range nums {
		newSet := make([][]int, 0)
		for _, c := range res {
			// []int{n} 相当于往前一个res结果的每一个子集内加入当前n 比如测试用例中的 [1] [2] [3]
			newSet = append(newSet, append([]int{n}, c...))
		}

		res = append(res, newSet...)
	}

	return res
}

// test cases [1,2,3]
// n 1
// c []
// set [[1]]
// res [[] [1]]
// n 2
// c []
// c [1]
// set [[2] [2 1]]
// res [[] [1] [2] [2 1]]
// n 3
// c []
// c [1]
// c [2]
// c [2 1]
// set [[3] [3 1] [3 2] [3 2 1]]
// res [[] [1] [2] [2 1] [3] [3 1] [3 2] [3 2 1]]

// @lc code=end
