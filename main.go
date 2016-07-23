package main

import (
	"flag"
	"fmt"

	"github.com/fatih/color"
	"github.com/ilopshire/eulerGo/problem31"
)

var problem int
var visual bool

func init() {
	flag.IntVar(&problem, "problem", 1, "Project Euler problem number")
	flag.BoolVar(&visual, "visual", false, "Show visual solution")
	flag.Parse()
}

func main() {

	color.Set(color.Bold)
	fmt.Printf("\n")
	fmt.Println("=========================================================")
	fmt.Printf("== Project Euler Problem #%d\n", problem)
	fmt.Println("=========================================================")
	fmt.Printf("\n")
	color.Unset()

	if visual {
		switch problem {
		case 31:
			problem31.SolveVisual()
		default:
			fmt.Println("Visual solution not found")
		}

	} else {
		switch problem {
		case 31:
			problem31.Solve()
		default:
			fmt.Println("Solution not found")
		}
	}

	color.Set(color.Bold)
	fmt.Printf("\n=========================================================\n\n")
	color.Unset()
}
