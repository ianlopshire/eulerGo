package problem2

import (
	"time"

	"github.com/ilopshire/eulerGo/util"
)

func Solve() {
	start := time.Now()

	answer := rec(0, 1, 1)

	util.PrintSolution(answer, time.Since(start))
}

func rec(sum, x, y int) int {
	if x+y < 4000000 {
		if (x+y)%2 == 0 {
			return rec(sum+x+y, y, x+y)
		}
		return rec(sum, y, x+y)

	}
	return sum
}
