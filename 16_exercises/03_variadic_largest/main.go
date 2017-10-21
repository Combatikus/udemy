package main

import "fmt"

func max(x ...int) int {
	var max int
	
	for _, i := range x {
		if i > max {
			max = i
		}
	}
	
	return max
}

func main() {
	fmt.Println(max(1, 2, 3, 4, 5))
}
