package main

func MaxOf2(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func MinOf2(i, j int) int {
	if i > j {
		return j
	} else {
		return i
	}
}

func abs(i int) int {
	if i < 0 {
		return - i
	} else {
		return i
	}
}