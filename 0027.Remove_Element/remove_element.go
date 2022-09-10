package _026_Remove_Duplicates_from_Sorted_Array

func removeElement(nums []int, val int) int {
	left := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] != val {
			nums[left] = nums[right]
			left++
		}
	}
	return left
}
