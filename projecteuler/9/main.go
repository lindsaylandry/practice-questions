/*
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a^2 + b^2 = c^2
For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/
package main

import (
	"fmt"
)

func main() {
	as := 0
	bs := 0
	cs := 0

	found := false
	for a := 1; a < 332; a++ {
		for b := a+1; b < 333; b++ {
			for c := 1000 - (b + a); c > 334; c-- {
				bb := 1000 - c - a
				if (a*a + bb*bb) == (c*c) {
					found = true
					cs = c
					bs = bb
					as = a
					break
				}
				if found {
					break
				}
			}
			if found {
				break
			}
		}
	}

	abc := as * bs * cs
	fmt.Printf("a: %d b: %d c: %d\n", as, bs, cs)
	fmt.Println(abc)
}
