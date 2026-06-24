func missingNumber(nums []int) int {
	maxSum, numSum := len(nums), 0
	for index, value := range nums {
		maxSum = index + maxSum
		numSum = value + numSum
	}
	return maxSum - numSum
}