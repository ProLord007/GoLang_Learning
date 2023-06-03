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
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Enter 5 integers:")
	arr = append(arr, 2, 3, 4, 5, 6, 7, 989, 0)
	fmt.Println("Array elements are:")
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
