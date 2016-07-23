package problem31

import (
	"bytes"
	"fmt"
	"time"

	"github.com/ilopshire/eulerGo/util"
)

func Solve() {
	start := time.Now()

	denominations := []int{1, 2, 5, 10, 20, 50, 100, 200}
	amount := 200
	cache := make([]int, amount+1)

	for _, d := range denominations {
		if d > amount {
			continue
		}

		cache[d] = cache[d] + 1
		for i := d; i <= amount; i++ {
			cache[i] = cache[i] + cache[i-d]
		}
	}

	util.PrintSolution(cache[amount], time.Since(start))
}

func SolveVisual() {
	start := time.Now()

	denominations := []int{1, 2, 5, 10, 20, 50, 100, 200}
	amount := 20
	cache := make([]int, amount+1)

	for _, d := range denominations {
		if d > amount {
			continue
		}

		cache[d] = cache[d] + 1
		for i := d; i <= amount; i++ {
			cache[i] = cache[i] + cache[i-d]
			printCahce(cache, d, i)
			time.Sleep(time.Millisecond * 100)
		}
	}

	fmt.Println("")
	util.PrintSolution(cache[amount], time.Since(start))
}

func printCahce(cache []int, d, index int) {
	var buffer bytes.Buffer
	// buffer.WriteString("\x1b[32;1mhello\x1b[32;1m lol")

	for i, v := range cache {
		if index == i {
			buffer.WriteString("\x1b[32;1m\x1b[7;1m")
		}

		if index-d == i {
			buffer.WriteString("\x1b[31;1m\x1b[7;1m")
		}
		buffer.WriteString(fmt.Sprintf("%3d", v))

		if index == i || index-d == i {
			buffer.WriteString("\x1b[0m")
		}

	}

	fmt.Printf("\r%2d [%s]\n", d, buffer.String())
}
