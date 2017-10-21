package main

import "fmt"

func foo(x ...int) {
	var max int

	for _, i := range x {
		if i > max {
			max = i
		}
	}

	fmt.Println(max)
}

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}
