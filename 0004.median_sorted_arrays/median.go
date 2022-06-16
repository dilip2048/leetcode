package _004_median_sorted_arrays

// this method will find the median of two sorted array
func findMedian(nums1, nums2 []int) float64 {
	x := len(nums1) + len(nums2)
	c := make([]int, x)
	m := 0
	n := 0
	i := 0
	for m < len(nums1) && n < len(nums2) {
		if nums1[m] < nums2[n] {
			c[i] = nums1[m]
			m++
		} else if nums1[m] > nums2[n] {
			c[i] = nums2[n]
			n++
		} else {
			c[i] = nums1[m]
			m++
		}
		i++
	}
	for m < len(nums1) {
		c[i] = nums1[m]
		i++
		m++
	}
	for n < len(nums2) {
		c[i] = nums2[n]
		i++
		n++
	}
	var median float64
	if x%2 == 1 {
		median = float64(c[(x-1)/2])
	} else {
		median = float64(c[(x-1)/2]+c[x/2]) / float64(2)
	}
	return median
}

/*
// A solution from a guy in leetcode in 5 lines.
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	if len(nums1)%2 == 0 {
		return (float64(nums1[len(nums1)/2]) + float64(nums1[len(nums1)/2-1])) / 2
	}
	return float64(nums1[len(nums1)/2])
}
*/
