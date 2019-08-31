package main

import "fmt"

func main() {
	x := 15
	y := 5
	z := chengfa(x, y)
	fmt.Printf("z=%d\n", z)
}

func jiafa(a, b int) int {
	return a + b
}

func jianfa(a, b int) int {
	return a - b
}

func chufa(a, b int) int {
	return a / b
}

func chengfa(a, b int) int {
	return a * b
}
