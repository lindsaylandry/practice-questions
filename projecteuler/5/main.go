/*
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/
package main

import (
	"fmt"
)

func main() {
	isDiv := false
	a := 20
	n := 2 * a
	for !isDiv {
		d := true
		for i := a; i > 2; i-- {
			if n % i != 0 {
				d = false
			}
		}

		if d {
			isDiv = true
		} else {
			n += a
		}
	}

	fmt.Println(n)
}
