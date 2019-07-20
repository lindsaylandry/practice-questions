/*
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/
package main

import (
	"fmt"
)

func main() {
	n := 1000
	a := 3
	b := 5

	nums := make(map[int]bool)

	for i := a; i < n; i += a {
		nums[i] = true
	}

	for i := b; i < n; i += b {
		nums[i] = true
	}

	sum := 0
	for k, _ := range nums {
		fmt.Println(k)
		sum += k
	}

	fmt.Println(sum)
}
