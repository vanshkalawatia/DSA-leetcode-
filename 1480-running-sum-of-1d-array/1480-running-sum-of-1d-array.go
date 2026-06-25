func runningSum(nums []int) []int {
    for i, num := range nums {
        if (i > 0) {
            nums[i] = num + nums[i-1]
        }
    }

    return nums
}