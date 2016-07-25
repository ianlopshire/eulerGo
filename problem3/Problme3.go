package problem3

import (
	"math"
	"time"

	"github.com/ilopshire/eulerGo/util"
)

func Solve() {
	start := time.Now()

	num := 600851475143
	var solution int

	max := int(math.Floor(math.Sqrt(float64(num))))
	for i := max; i >= 2; i-- {
		if num%i == 0 && isPrime(i) {
			solution = i
			break
		}
	}

	util.PrintSolution(solution, time.Since(start))
}

func isPrime(num int) bool {
	max := int(math.Floor(math.Sqrt(float64(num))))

	for i := 2; i <= max; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
