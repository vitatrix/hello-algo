package chapter_sorting

func bubbleSort(nums []int) {
	// 迭代待排区间("水面 i")
	for i := len(nums); i > 0; i-- {
		// 相邻元素比较，大的不断右移(交换位置)
		for j := 1; j < i; j++ {
			if nums[j - 1] > nums[j] {
				nums[j - 1], nums[j] = nums[j], nums[j - 1]
			}
		}
	}
}

func bubbleSortWithFlag(nums []int) {
	// 迭代待排区间
	for i := len(nums); i > 0; i-- {
		k := true
		for j := 1; j < i; j++ {
			for nums[j - 1] > nums[j] {
				nums[j - 1], nums[j], k = nums[j], nums[j - 1], false
			}
		}
		// 没有交换说明已经排好，提前结束
		if k {
			break
			// return
		}
	}
}
