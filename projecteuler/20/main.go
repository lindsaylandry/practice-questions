/*
n! means n × (n − 1) × ... × 3 × 2 × 1

For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!
*/

package main

import (
	"os"
	"fmt"
	"strconv"
	"math/big"
)

func main() {
	num := 100
	if len(os.Args) >= 2 {
		var err error
		num, err = strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
	}

	fact := getFactorial(num)

	fmt.Printf("%d! = %d\n", num, fact)

	str := fact.Text(10)

	a := 0
	for _, c := range(str) {
		a += int(c - '0')
	}

	fmt.Println(a)
}

func getFactorial(n int) *big.Int {
	a := big.NewInt(1)
	for i := 1; i <= n; i++ {
		a = a.Mul(a, big.NewInt(int64(i)))
	}

	return a
}
