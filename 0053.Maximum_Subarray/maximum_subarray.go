package _053_Maximum_Subarray

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	currSum := 0
	for i := 0; i < len(nums); i++ {
		// it won't contribute in finding the max array
		if currSum < 0 {
			currSum = 0
		}
		currSum += nums[i]
		if currSum > maxSum {
			maxSum = currSum
		}
	}
	return maxSum
}
