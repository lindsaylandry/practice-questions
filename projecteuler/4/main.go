/*
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b int

	p := 0
	for i := 999; i > 1; i-- {
		for j := i; j > 1; j-- {
			num := i*j
			if isPalindrome(num) {
				if num > p {
					p = num
					a = i
					b = j
				}
				break
			}
		}
	}

	fmt.Printf("Palindrome: %d (%d x %d)\n", p, a, b)
}

func isPalindrome(n int) bool {
	str := strconv.Itoa(n)

	p := true
	for i, _ := range str {
		if str[i] != str[len(str)-1-i] {
			p = false
			break
		}
	}

	return p
}
