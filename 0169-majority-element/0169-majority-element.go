func majorityElement(nums []int) int {
	maxRepeat := map[int]int{}

	for _, value := range nums {
		
		maxRepeat[value]++
	}

	for key, value := range maxRepeat {
		if value > len(nums)/2 {
			return key
		}
	}
	return 0
}