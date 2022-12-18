package main

/*
 * @lc app=leetcode id=786 lang=golang
 *
 * [786] K-th Smallest Prime Fraction
 */

// @lc code=start
func kthSmallestPrimeFraction(arr []int, k int) []int {
	//因为是真分数的范围在[0,1]之间
	//所以初始化下界0 上界1
	lo, hi, n := 0.0, 1.0, len(arr)
	for {
		mid, cnt, p, q, j := (lo+hi)/2.0, 0, 0, 1, 0
		for i := 0; i < n; i++ {
			// 统计整个数组内比mid大的所有分数的数目
			// arr[i]/arr[j] > mid
			for j < n && float64(arr[i]) > float64(mid)*float64(arr[j]) {
				j++
			}
			// 比 mid小的分数
			cnt += n - j
			// 记录p,q的最小值
			// q/p > arr[j]/arr[i]
			if j < n && q*arr[i] > p*arr[j] {
				p = arr[i]
				q = arr[j]
			}
		}
		if cnt == k {
			return []int{p, q}
		} else if cnt < k {
			lo = mid
		} else {
			hi = mid
		}

	}

}

// @lc code=end
