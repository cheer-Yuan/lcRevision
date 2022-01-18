package main

func findMinDifference(timePoints []string) int {
	leng := len(timePoints)
	if leng > 60 * 24 {
		return 0
	}

	hash := map[int]bool{}

	for _, i := range timePoints {
		key := converse(i)
		if hash[key] == true {
			return 0
		} else {
			hash[key] = true
		}
	}

	result, temp := 60 * 24, 0
	for i := 0; i < 60 * 24; i++ {
		if hash[i] == true {
			if temp == 0 {
				temp = i
				continue
			}

			if i - temp < result {
				result = i - temp
			}

			temp = i
		}
	}

	return result
}

func converse(time string) int {
	b1, b2, b3, b4 := time[0] - '0', time[1] - '0', time[3] - '0', time[4] - '0'

	return int(b1) * 10 * 60 + int(b2) * 60 + int(b3) * 10 + int(b4)
}