/*
given a sequence of numbers, return all permutations of strings that correspond with
the numbers on a keypad
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	d := 2451234

	s, err := strings.Itoa(d)
	if err != nil {
		panic(err)
	}

	str := permute(s)
	fmt.Println(str)
}

func permute(str []string) []string {
	if len(str) == 0 {
		return []string{str}
	}
	
	t := []string{}
	for i, p := range str {
		t = permute(str[i:])
		
	}
}
