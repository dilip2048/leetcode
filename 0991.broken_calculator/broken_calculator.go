package _991_broken_calculator

func brokenCalc(startValue int, target int) int {
	counter := 0
	for target != startValue {
		if target == startValue {
			return counter + 1
		} else if target%2 == 0 && target > startValue {
			target /= 2
		} else {
			target += 1
		}
		counter++
	}
	return counter
}
