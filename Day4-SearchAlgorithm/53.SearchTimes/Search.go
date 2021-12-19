func search(nums []int, target int) int {
	r1 := rightIndex(nums, target)
	r2 := rightIndex(nums, target-1)
	return r1 - r2
}

func rightIndex(n []int, t int) int {
	i := 0
	j := len(n) - 1
	for i <= j {
		m := (i + j) / 2
		if t < n[m] {
			j = m - 1
		} else {
			i = m + 1
		}
	}
	return i
}