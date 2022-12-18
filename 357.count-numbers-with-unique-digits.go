/*
 * @lc app=leetcode id=357 lang=golang
 *
 * [357] Count Numbers with Unique Digits
 */

// @lc code=start
func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	res, uniqDigits, availableNums := 10, 9, 9
	for n > 1 && availableNums > 0 {
		uniqDigits *= availableNums
		res += uniqDigits
		availableNums--
		n--
	}
	return res
}

// @lc code=end

