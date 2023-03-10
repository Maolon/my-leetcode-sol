package main

import "math"

/*
 * @lc app=leetcode id=509 lang=golang
 *
 * [509] Fibonacci Number
 */

// @lc code=start
func fib(n int) int {
	if n <= 1 {
		return n
	}
	fib := make([]int, n+1)
	fib[0], fib[1] = 0, 1

	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib[n]
}

// 还有一些数学解法

// 矩阵快速幂 时间复杂度 O(log n)，空间复杂度 O(log n)
// | 1 1 | ^ n   = | F(n+1) F(n)   |
// | 1 0 |		   | F(n)	F(n-1) |
func fib4(N int) int {
	if N <= 1 {
		return N
	}
	var A = [2][2]int{
		{1, 1},
		{1, 0},
	}
	A = matrixPower(A, N-1)
	return A[0][0]
}

func matrixPower(A [2][2]int, N int) [2][2]int {
	if N <= 1 {
		return A
	}
	A = matrixPower(A, N/2)
	A = multiply(A, A)

	var B = [2][2]int{
		{1, 1},
		{1, 0},
	}
	if N%2 != 0 {
		A = multiply(A, B)
	}

	return A
}

func multiply(A [2][2]int, B [2][2]int) [2][2]int {
	x := A[0][0]*B[0][0] + A[0][1]*B[1][0]
	y := A[0][0]*B[0][1] + A[0][1]*B[1][1]
	z := A[1][0]*B[0][0] + A[1][1]*B[1][0]
	w := A[1][0]*B[0][1] + A[1][1]*B[1][1]
	A[0][0] = x
	A[0][1] = y
	A[1][0] = z
	A[1][1] = w
	return A
}

// 公式法 f(n)=(1/√5)*{[(1+√5)/2]^n -[(1-√5)/2]^n}，用 时间复杂度在 O(log n) 和 O(n) 之间，空间复杂度 O(1)
// 经过实际测试，会发现 pow() 系统函数比快速幂慢，说明 pow() 比 O(log n) 慢
// 斐波那契数列是一个自然数的数列，通项公式却是用无理数来表达的。而且当 n 趋向于无穷大时，前一项与后一项的比值越来越逼近黄金分割 0.618（或者说后一项与前一项的比值小数部分越来越逼近 0.618）。
// 斐波那契数列用计算机计算的时候可以直接用四舍五入函数 Round 来计算。
func fib5(N int) int {
	var goldenRatio float64 = float64((1 + math.Sqrt(5)) / 2)
	return int(math.Round(math.Pow(goldenRatio, float64(N)) / math.Sqrt(5)))
}

// 另外还有一些胡里花哨的解法

// @lc code=end
