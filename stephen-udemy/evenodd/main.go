package main

import "fmt"

func main() {
	num := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, v := range num {
		if v%2 == 1 {
			fmt.Println("Number is odd ", v)
		} else {
			fmt.Println("Number is even ", v)
		}
	}
}