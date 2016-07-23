package util

import (
	"fmt"
	"time"
)

func PrintSolution(s interface{}, d time.Duration) {
	fmt.Printf("Solution: %+v\nExecution Time: %s\n", s, d)
}
