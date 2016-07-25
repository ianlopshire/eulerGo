package util

import "math"

func IsPrime(num int) bool {
	max := int(math.Floor(math.Sqrt(float64(num))))

	for i := 2; i <= max; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
