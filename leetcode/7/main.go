/*
Given a 32-bit signed integer, reverse digits of an integer.

Example 1:

Input: 123
Output: 321
Example 2:

Input: -123
Output: -321
Example 3:

Input: 120
Output: 21
Note:
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
*/
package main

import (

)

func reverse(x int) int {
    rev := 0
    isNeg := false
    if x < 0 {
        isNeg = true
        x = -x
    }
    
    for x > 0 {
        rev += x % 10
        rev = rev * 10
        x = x / 10
    }
    
    rev = rev / 10
    
    if rev > 2147483647 {
        return 0
    }
    
    if isNeg {
        rev = -rev
    }
    
    return rev
}
