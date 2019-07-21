/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/
package main

import (
	"fmt"
	"os"
	"strconv"
	"math"
)

func main() {
	n := 600851475143

	if len(os.Args) >= 2 {
		var err error
		n, err = strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
	}

	lPrime := 0
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n % i == 0 {
			if isPrime(i) {
				lPrime = i
			}
		}
	}

	fmt.Printf("%d prime: %d\n", n, lPrime)
}

func isPrime(n int) bool {
	// find mod of all nums up to sqrt n
	c := int(math.Sqrt(float64(n)))

	isPrime := true
	for i := 2; i < c; i++ {
		if n % i == 0 {
			isPrime = false
			break
		}
	}

	return isPrime
}
