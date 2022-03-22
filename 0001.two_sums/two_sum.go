package twosum

// this method will find the two indexes whose elements' sum is equal to the target number
func twoSum(nums []int, target int) []int {
	//map to keep track of visited elements
	h := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if index, ok := h[target-nums[i]]; ok {
			return []int{index, i}
		}
		// mark the elements visited and save the index in the map
		h[nums[i]] = i
	}
	return nil
}
