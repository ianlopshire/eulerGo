package util

func LCM(a, b int) int {
	return a * b / GDC(a, b)
}
