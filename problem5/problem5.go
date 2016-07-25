package problem5

import (
	"time"

	"github.com/ilopshire/eulerGo/util"
)

func Solve() {
	start := time.Now()

	min := 1
	max := 20

	temp := max

	for i := min; i < max; i++ {
		temp = util.LCM(temp, i)
	}

	util.PrintSolution(temp, time.Since(start))
}
