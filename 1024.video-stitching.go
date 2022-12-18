package main

import "sort"

/*
 * @lc app=leetcode id=1024 lang=golang
 *
 * [1024] Video Stitching
 */

// @lc code=start
func videoStitching(clips [][]int, time int) int {
	// 贪心算法
	// 重复区域
	if time == 0 {
		return 0
	}
	// 按照start升序，如果start相同按end排序
	sort.Slice(clips, func(x, y int) bool {
		if clips[x][0] == clips[y][0] {
			return clips[x][1] < clips[y][1]
		} else {
			return clips[x][0] < clips[y][0]
		}
	})
	res, curEnd, nextEnd, i := 0, 0, 0, 0
	for i < len(clips) && clips[i][0] <= curEnd {
		// 贪心选择到下一个视频
		for i < len(clips) && clips[i][0] <= curEnd {
			nextEnd = max(nextEnd, clips[i][1])
			i++
		}
		res++
		// 更新curEnd
		curEnd = nextEnd
		if curEnd >= time {
			return res
		}
	}
	//没找到
	return -1

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
