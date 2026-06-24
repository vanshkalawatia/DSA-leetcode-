func singleNumber(nums []int) int {
    var num int
    for _,value := range(nums){
        num ^= value
    }
    return num
}