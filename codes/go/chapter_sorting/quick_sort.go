package chapter_sorting

type quickSort struct{}

func (q *quickSort) partition(nums []int, left, right int) int {
	i, j := left, right
	for i < j {
		for j > i && nums[j] >= nums[left] {
			j--
		}
		for i < j && nums[i] <= nums[left] {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[left], nums[i] = nums[i], nums[left]
	return i
}

func (q *quickSort) quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	pivot := q.partition(nums, left, right)
	q.quickSort(nums, left, pivot - 1)
	q.quickSort(nums, pivot + 1, right)
}

type quickSortMedian struct{}

func (q *quickSortMedian) mediThree(nums []int, left, mid, right int) int {
	// 异或逻辑
	if (nums[mid] < nums[left]) != (nums[mid] < nums[right]) {
		return mid
	} else if (nums[right] < nums[mid]) != (nums[right] < nums[left]) {
		return right
	}
	return left
}

func (q *quickSortMedian) partition(nums []int, left, right int) int {
	// 中位数
	mid := q.mediThree(nums, left, (left + right)/2, right)
	// 中位数交换到最左边
	nums[left], nums[mid] = nums[mid], nums[left]
	// 分区
	i, j := left, right
	for i < j {
		for j > i && nums[j] >= nums[left] {
			j--
		}
		for i < j && nums[i] <= nums[left] {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[left], nums[i] = nums[i], nums[left]
	return i
}

func (q *quickSortMedian) quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	pivot := q.partition(nums, left, right)
	q.quickSort(nums, left, pivot - 1)
	q.quickSort(nums, pivot + 1, right)
}

type quickSortTailCall struct{}

func (q *quickSortTailCall) partition(nums []int, left, right int) int {
	i, j := left, right
	for i < j {
		for j > i && nums[j] >= nums[left] {
			j--
		}
		for i < j && nums[i] <= nums[left] {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[left], nums[i] = nums[i], nums[left]
	return i
}

func (q *quickSortTailCall) quickSort(nums []int, left, right int) {
	for left < right {
		pivot := q.partition(nums, left, right)
		if (pivot - left) < (pivot - right) {
			q.quickSort(nums, left, pivot - 1)
			left = pivot + 1
		} else {
			q.quickSort(nums, pivot + 1, right)
			right = pivot - 1
		}
	}
}