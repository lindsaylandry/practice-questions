/*
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.
*/
package main

import (
	"fmt"
)

func isValid(s string) bool {
    fmt.Println(s)

		// keep stack of all open parens, drain when finding closing pairs
		openStack := ""

    for _, i := range s {

        // push onto stack
        if i == '(' || i == '[' || i == '{' {
            openStack = openStack + string(i)
        } else {
            if len(openStack) == 0 {
                return false
            }
            if i == ')' {
                if openStack[len(openStack)-1] != '(' {
                    return false
                } else {
                    // pop stack
                    openStack = openStack[0:len(openStack)-1]
                }
            } else if i == ']' {
                if openStack[len(openStack)-1] != '[' {
                    return false
                } else {
                    // pop stack
                    openStack = openStack[0:len(openStack)-1]
                }
            } else if i == '}' {
                if openStack[len(openStack)-1] != '{' {
                    return false
                } else {
                    // pop stack
                    openStack = openStack[0:len(openStack)-1]
                }
            }
        }
    }

    // stack must be empty after translating thru string
    if len(openStack) != 0 {
        return false
    }

    return true
}
