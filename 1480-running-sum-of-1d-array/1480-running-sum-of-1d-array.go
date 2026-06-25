func runningSum(nums []int) []int {
    if len(nums) < 2 {
        return nums
    }
 
    for i,num := range nums {
        if i > 0 {
        nums[i] = num + nums[i-1]
        }
    }
 
 
 
    return nums
}