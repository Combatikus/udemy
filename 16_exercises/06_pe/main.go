package main

import "fmt"

func fib(x int) int {
	if x <= 2 {
		return 1
	}

	return fib(x-1) + fib(x-2)
}

func main() {
	/*var sum int

	for i := 0; i < 1001; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	fmt.Println(sum)*/

	fmt.Println(fib(10))
}
