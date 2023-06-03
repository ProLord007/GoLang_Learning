package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	/*	file, err := os.Open("test.txt")
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
		fmt.Println(str)*/
	iobs, err := ioutil.ReadFile("test.txt")
	if err == nil {
		fmt.Println(string(iobs))
	} else {
		fmt.Println(err)
	}
}
