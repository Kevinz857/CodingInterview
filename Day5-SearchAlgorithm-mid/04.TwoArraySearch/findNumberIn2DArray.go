
/*暴力查找*/
func findNumberIn2DArray(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}

/*
二分法查找
func findNumberIn2DArray(matrix [][]int, target int) bool{
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	for i := 0; i < len(matrix); i++ {
		l := 0
		r := len(matrix[i]) -1
		for l <= r {
			mid := l + (r - l ) /2
			if matrix[i][mid] < target {
				l = mid + 1
			}else if matrix[i][mid] > target {
				r = mid -1
			} else {
				return true
			}
		}
	}
	return false
}
*/