package chapter_sorting

func selectionSort(nums []int) {
	n := len(nums)
	// 待排区间 [i, n)
	for i := 0; i < n; i++ {
		k := i
		// 线性查找最小值
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[k] {
				// 记录最小值的位置
				k = j
			}
		}
		// 交换位置
		nums[i], nums[k] = nums[k], nums[i]
		// if k != i {
		// 	nums[i], nums[k] = nums[k], nums[i]
		// }
	}
}
