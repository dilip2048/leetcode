package _026_Remove_Duplicates_from_Sorted_Array

// we will use two pointer system
// right pointer will point all the unique element
// left pointer will point to the last unique element
func removeDuplicates(nums []int) int {
	left := 0
	for right := 1; right < len(nums); right++ {
		//skip and find the unique element
		if nums[left] != nums[right] {
			left++
			// increment left and put the unique elements to the left index
			nums[left] = nums[right]
		}
	}
	return left + 1
}
