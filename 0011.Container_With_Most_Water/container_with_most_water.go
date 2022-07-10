package _011_Container_With_Most_Water

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	var maxArea int
	for left < right {
		b := right - left
		l := height[left]
		if height[left] > height[right] {
			l = height[right]
		}
		area := b * l
		if area > maxArea {
			maxArea = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}
