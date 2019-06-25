package main

import "fmt"

func main() {

	a := [6]int{1, 2, 3, 4, 5, 6}
	var temp int

	i := 0
	j := len(a) - 1

	for i < j {
		temp = a[i]
		a[i] = a[j]
		a[j] = temp
		i++
		j--
	}

	i = 0
	for i < 6 {
		fmt.Println(a[i])
		i++
	}

	// Just testing for loop
	for i := 0; i < 5; i++ {
		fmt.Printf("Great\n")
	}

}
