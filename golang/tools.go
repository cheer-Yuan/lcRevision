package main

func MaxOf2(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func abs(i int) int {
	if i < 0 {
		return - i
	} else {
		return i
	}
}