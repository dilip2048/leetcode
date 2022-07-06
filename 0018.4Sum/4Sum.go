package _018_4Sum

import "sort"

func fourSum(nums []int, target int) [][]int {
	var result [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		if i == 0 || i > 0 && nums[i] != nums[i-1] {
			// if this is true then there will be no further quadruple which equals target.
			if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
				break
			}
			// this will take us to nearest possible quadruple
			if nums[i]+nums[len(nums)-3]+nums[len(nums)-2]+nums[len(nums)-1] < target {
				continue
			}
			for j := i + 1; j < len(nums)-2; j++ {
				if j == i+1 || (j > i+1 && nums[j] != nums[j-1]) {
					if nums[i]+nums[j] > target-nums[j+1]-nums[j+2] {
						break
					}
					if nums[i]+nums[j] < target-nums[len(nums)-2]-nums[len(nums)-1] {
						continue
					}

					low := j + 1
					high := len(nums) - 1
					for low < high {
						if nums[i]+nums[j]+nums[low]+nums[high] == target {
							result = append(result, []int{nums[i], nums[j], nums[low], nums[high]})
							//avoid duplicates
							for low < high && nums[low] == nums[low+1] {
								low++
							}
							for low < high && nums[high] == nums[high-1] {
								high--
							}
							low++
							high--
						} else if nums[i]+nums[j]+nums[low]+nums[high] < target {
							low++
						} else {
							high--
						}
					}
				}
			}
		}
	}
	return result
}
