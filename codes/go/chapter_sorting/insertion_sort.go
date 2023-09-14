package chapter_sorting

func insertionSort(nums []int) {
	// 迭代已排区间
	n := len(nums)
	for i := 0; i < n - 1; i++ {
		s := nums[i + 1]
		// 从右往左插入
		j := i
		for j >= 0 && nums[j] > s {
			nums[j + 1] = nums[j]
			j--
		}
		// j + 1 是腾出来的位置
		nums[j + 1] = s
	}
}

// func insertionSort(nums []int) {
// 	// 迭代已排区间
// 	n := len(nums)
// 	for i := 0; i < n - 1; i++ {
// 		s := nums[i + 1]
// 		// 插入已排区间 第一直觉 从左往右插
// 		for j := 0; j <= i + 1; j++ {
// 			if s < nums[j] {
// 				// [j, i] 都右移一个位置
// 				for k := i + 1; k > j; k-- {
// 					nums[k] = nums[k - 1]
// 				}
// 				nums[j] = s
// 				break
// 			}
// 		}
// 	}
// }