package _016_3Sum_Closest

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	minDiff := int(math.MaxInt32)
	var resultSum int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		low := i + 1
		high := len(nums) - 1
		for low < high {
			sum := nums[i] + nums[low] + nums[high]
			if sum > target {
				high--
			} else {
				low++
			}
			diff := sum - target
			if diff < 0 {
				diff = -1 * diff
			}
			if diff < minDiff {
				minDiff = diff
				resultSum = sum
			}
		}
	}
	return resultSum
}
