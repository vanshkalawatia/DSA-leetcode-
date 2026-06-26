func pivotIndex(nums []int) int {
	var total int
	var leftSum int

	for _, value := range nums {
		total += value
	}

	for i, n := range nums {

		if leftSum == total - leftSum - n {
			return i
		}
		leftSum += n
	}

	return -1
}