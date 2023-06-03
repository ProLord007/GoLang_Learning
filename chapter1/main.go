package main

import (
	"fmt"
	"math"
	"os"
)

type Circle struct {
	x, y, r float64
}

func (c Circle) power() float64 {
	c.r = 2
	return (c.r * c.r * math.Pi)
}
func main() {
	c := Circle{x: 0, y: 0, r: 5}
	file, err := os.Open("test.txt")
	if err != nil {
		return
	}
	defer file.Close()

	//get the file size
	stat, err := file.Stat()
	if err != nil {
		return
	}
	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)
	fmt.Println(c.power())
}
