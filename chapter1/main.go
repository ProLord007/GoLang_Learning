package main

import (
	"fmt"
	"math"
)

func f(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var b int = 5
	fmt.Scanf("%d", &b)
	fmt.Println(f(b))
}
