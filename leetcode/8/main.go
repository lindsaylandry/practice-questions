/*
Implement atoi which converts a string to an integer.

The function first discards as many whitespace characters as necessary until the first non-whitespace character is found. Then, starting from this character, takes an optional initial plus or minus sign followed by as many numerical digits as possible, and interprets them as a numerical value.

The string can contain additional characters after those that form the integral number, which are ignored and have no effect on the behavior of this function.

If the first sequence of non-whitespace characters in str is not a valid integral number, or if no such sequence exists because either str is empty or it contains only whitespace characters, no conversion is performed.

If no valid conversion could be performed, a zero value is returned.

Note:

Only the space character ' ' is considered as whitespace character.
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−2^31,  2^31 − 1]. If the numerical value is out of the range of representable values, INT_MAX (2^31 − 1) or INT_MIN (−2^31) is returned.
Example 1:

Input: "42"
Output: 42
Example 2:

Input: "   -42"
Output: -42
Explanation: The first non-whitespace character is '-', which is the minus sign.
             Then take as many numerical digits as possible, which gets 42.
Example 3:

Input: "4193 with words"
Output: 4193
Explanation: Conversion stops at digit '3' as the next character is not a numerical digit.
Example 4:

Input: "words and 987"
Output: 0
Explanation: The first non-whitespace character is 'w', which is not a numerical 
             digit or a +/- sign. Therefore no valid conversion could be performed.
Example 5:

Input: "-91283472332"
Output: -2147483648
Explanation: The number "-91283472332" is out of the range of a 32-bit signed integer.
             Thefore INT_MIN (−2^31) is returned.
*/
package main

import (
    "math"
)

func myAtoi(str string) int {
    // remove whitespace
    for i, a := range str {
        if a != ' ' {
            str = str[i:]
            break
        }
    }
    
    if len(str) == 0 {
        return 0
    }
    
    // find out if negative
    isNeg := false
    if str[0] == '-' {
        isNeg = true
        // absolute number
        str = str[1:]
    } else if str[0] == '+' {
        // skip + sign
        str = str[1:]
    }

    n := int32(0)
    
    max := false
    // iterate thru string and convert to int
    for _, a := range str {
        // break if rune is not a number (0-9)
        b := int(a - '0')
        if b < 0 || b > 9 {
            break
        }
        
        prev := n
        n = n * 10
        if prev > n || n % 10 > 0 {
            n = math.MaxInt32
            max = true
            break
        }
        
        // break if reached largest number        
        prev = n
        n += int32(b)
        if prev > n {
            n = math.MaxInt32
            max = true
            break
        }
    }
    
    if isNeg {
        if max {
            n = math.MinInt32
        } else {
            n = -n
        }
    }
    
    return int(n)
}
