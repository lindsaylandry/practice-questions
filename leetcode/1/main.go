/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/
package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	indices := make([]int, 2)
	found := false
	for a, i := range nums {
		for b, j := range nums {
			if a == b {
				continue
			} else if i + j == target {
				indices[0] = a
				indices[1] = b
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	return indices
}
