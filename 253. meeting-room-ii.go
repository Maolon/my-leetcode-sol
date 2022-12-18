package main

import "sort"

// 返回需要申请的会议室数量
// 时间重复问题
// 贪心算法
func minMeetingRooms(meetings [][]int) int {
	begin := make([]int, len(meetings)+1)
	end := make([]int, len(meetings)+1)
	sort.Ints(begin)
	sort.Ints(end)
	for i := range meetings {
		begin[i] = meetings[i][0]
		end[i] = meetings[i][1]
	}
	cnt, res, i, j := 0, 0, 0, 0
	for i < len(meetings) && j < len(meetings) {
		if begin[i] < end[j] {
			i++
			//扫描到一个红点
			cnt++
		} else {
			//扫描到一个绿点
			cnt--
			j++
		}
		res = max(res, cnt)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
