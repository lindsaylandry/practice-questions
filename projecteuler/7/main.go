/*
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	c := 1
	var a uint64 = 2

	for c < 10001 {
		a += 1
		if isPrime(a) {
			c += 1
		}
	}

	fmt.Println(a)
}

func isPrime(n uint64) bool {
	c := int(math.Sqrt(float64(n)))
	isPrime := true
	for i := 2; i <= c; i++ {
		if n % uint64(i) == 0 {
			isPrime = false
			break
		}
	}

	return isPrime
}
