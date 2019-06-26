package main

import "fmt"

func main() {
	slice := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}

	for i, element := range slice {
		fmt.Println(i, element)
	}
}
