package main

import (
	"fmt"
)

func main() {
	f := factorial(4)
	fact := 1
	for n := range f {
		fact *= n
	}

	fmt.Println("Total: ", fact)
}

func factorial(n int) <-chan int {
	out := make(chan int)
	go func() {
		for index := n; index > 0; index-- {
			out <- index
		}

		close(out)
	}()

	return out
}
