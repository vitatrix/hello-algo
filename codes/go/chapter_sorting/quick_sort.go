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
