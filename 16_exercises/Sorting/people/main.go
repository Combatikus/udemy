package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}

	fmt.Println(nums)

	fmt.Printf("%T\n", nums)
	other := sort.IntSlice(nums)
	fmt.Printf("%T\n", other)

	i := sort.Reverse(sort.IntSlice(nums))

	fmt.Println(i)
	fmt.Printf("%T\n", i)

	sort.Sort(i)
	fmt.Println(nums)
}
