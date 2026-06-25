func runningSum(nums []int) []int {
	runSum := []int{}
	var sum int

	for _, value := range nums {
		sum = value + sum
		runSum = append(runSum, sum)
	}
	return runSum
}
