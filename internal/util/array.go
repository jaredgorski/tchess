package array

import ()

func Contains(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}

	return false
}
