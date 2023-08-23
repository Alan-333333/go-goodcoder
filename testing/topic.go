package testing

func QuickSort(nums []int) []int {

	if len(nums) <= 1 {
		return nums
	}

	pivot := len(nums) / 2

	var low, hight []int

	for _, value := range nums {
		if value < nums[pivot] {
			low = append(low, value)
		} else if value > nums[pivot] {
			hight = append(hight, value)
		}
	}

	low = QuickSort(low)
	hight = QuickSort(hight)

	result := append(low, nums[pivot])
	result = append(result, hight...)
	return result
}
