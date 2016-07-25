package problem1

import (
	"time"

	"github.com/ilopshire/eulerGo/util"
)

func Solve() {
	start := time.Now()
	answer := 0

	for i := 1; i < 1000; i++ {
		if i%5 == 0 || i%3 == 0 {
			answer += i
		}
	}

	util.PrintSolution(answer, time.Since(start))
}
