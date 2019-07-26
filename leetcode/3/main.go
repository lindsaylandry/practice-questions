/*
Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3 
Explanation: The answer is "abc", with the length of 3. 
Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3. 
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/
package main

import (
	"fmt"
)

func main() {
	num := lengthOfLongestSubstring("abcabcbb")
	fmt.Println(num)
}

func lengthOfLongestSubstring(s string) int {
    longest := 0
    last := 0
    isThere := make(map[rune]bool)
    for i, _ := range s {
        if len(s[i:]) < longest {
            break
        }
        for _, r := range s[i+last:] {
            if _, ok := isThere[r]; !ok {
                isThere[r] = true
                last += 1
            } else {
                break
            }
        }
        if last > longest {
            longest = last
        }
        delete(isThere, rune(s[i]))
        last -= 1
    }

    return longest
}
