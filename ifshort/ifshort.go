package main

import "fmt"

func main() {
	var c = 2
	if v := c * c; v < 10 {
		fmt.Println(v)
	}
}
