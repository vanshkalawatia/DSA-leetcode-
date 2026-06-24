func singleNumber(nums []int) int {
    single := map[int] bool{}
    for _,value := range nums{
        if !single[value]{
            single[value] = true
        }else {
            delete(single,value)
        }
    }
    for key ,value := range single {
        if value == true {
            return key
        }
    }
    return 0
}