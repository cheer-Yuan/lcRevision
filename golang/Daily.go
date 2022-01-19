package main

/*539. 最小时间差
给定一个 24 小时制（小时:分钟 "HH:MM"）的时间列表，找出列表中任意两个时间的最小时间差并以分钟数表示。
*/
func findMinDifference(timePoints []string) int {
	leng := len(timePoints)
	if leng > 60 * 24 {     // 必有重复时间
		return 0
	}

	hash := map[int]bool{}

	var converse func(time string) int		// 时间转换
	converse = func(time string) int {
		b1, b2, b3, b4 := time[0] - '0', time[1] - '0', time[3] - '0', time[4] - '0'

		return int(b1) * 10 * 60 + int(b2) * 60 + int(b3) * 10 + int(b4)
	}

	for _, i := range timePoints {  // 初始化hash表
		key := converse(i)
		if hash[key] == true {
			return 0
		} else {
			hash[key] = true
		}
	}

	result, temp, earliestTime := 1440, 0, 1440      // 初始化最小时间间隔
	for i := 0; i < 1440; i++ {
		if hash[i] == true {

			if result == 1440 {     // 第一个时间
				temp = i
				result--
				earliestTime = i
				continue
			}

			if earliestTime > i {   // 更新最早的时间
				earliestTime = i
			}
			difference := i - temp  // 用最早时间计算环形时间差
			difference2 := 1440 - i + earliestTime
			if difference2 < difference {
				difference = difference2
			}

			if difference < result {
				result = difference
			}

			temp = i
		}
	}

	return result
}