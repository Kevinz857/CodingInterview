func findRepeatNumber(nums []int) int {
	maps := make(map[int]bool)
	for _, value := range nums {
		if maps[value] {
			return false
		} else {
			maps[value] = true
		}

	}
	return -1
}

/*
数组中的数字都在0~n-1的范围内，如果这个数组中没有重复数字，那么当数组排序之后数字i将出现在下标为i的位置。由于数组中有重复的数字，有些位置可能存在多个数字。有些位置没有数字。
现在让我们重排这个数组，当扫描到下表为i的数字时，首先比较这个数字（用n表示）是不是等于i 。如果是，则继续扫描下一个数字，如果不是，则再拿它喝第n个数字进行比较，如果它和第n个数字相等，则找到了重复的数字（这个数字和下标为i和n的位置都出现了） ，如果它和第n个数字不相等，就把第i个数字和第n个数字交换，把数字n放在对应下标自己的位置。然后重复比较，直到发现这个重复的数字。
func findRepeatNumber(nums []int) int{
	for index, value := range nums {
		if index == value {
			continue
		}
		if nums[value] == value {
			return value
		}
		nums[value], nums[index] = nums[index], nums[value]
	}
	return -1
}
*/

/*
将数组排序，遍历有序数组，只要相邻两个元素相同，就有相同元素，输出这个元素

func findRepeatNumber(nums []int) int  {
	sort.Ints(nums)
	for i, _ := range nums {
		if i > 0 {
			if nums[idx] == nums[idx-1]{
				return nums[idx]
			}
		}
	}
	return -1
}
*/