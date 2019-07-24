/*
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	top := 2000000
	sum := uint64(2)

	for i := 3; i < top; i++ {
		if isPrime(i) {
			sum = sum + uint64(i)
		}
	}

	fmt.Println(sum)
}

func isPrime(n int) bool {
	c := int(math.Sqrt(float64(n)))
	p := true
	for i := 2; i <= c; i++ {
		if n % i == 0 {
			p = false
			break
		}
	}

	return p
}
