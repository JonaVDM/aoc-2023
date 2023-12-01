package utils

func Max(arr []int) int {
	max := arr[0]

	for _, i := range arr {
		if i > max {
			max = i
		}
	}

	return max
}

func Min(arr []int) int {
	min := arr[0]

	for _, i := range arr {
		if i < min {
			min = i
		}
	}

	return min
}

func IntRange(start, stop int) []int {
	out := make([]int, 0)
	for i := start; i < stop; i++ {
		out = append(out, i)
	}
	return out
}

func Overlap(a, b []int) bool {
	for _, x := range a {
		for _, y := range b {
			if x == y {
				return true
			}
		}
	}

	return false
}
