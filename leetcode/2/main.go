/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.
*/
package main

import (
	"fmt"
	"strconv"
	"math/big"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	n1 := ListNode {
		Val: 7,
		Next: nil,
	}
	n2 := ListNode {
		Val: 4,
		Next: &n1,
	}
	n3 := ListNode {
		Val: 5,
		Next: &n2,
	}

	m1 := ListNode {
    Val: 9,
    Next: nil,
  }
  m2 := ListNode {
    Val: 1,
    Next: &m1,
  }
  m3 := ListNode {
    Val: 8,
    Next: &m2,
  }

	sum := addTwoNumbers(&n3, &m3)

	for sum != nil {
		fmt.Println(sum.Val)
		sum = sum.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	lll := ""
	for l1 != nil {
		lll = strconv.Itoa(l1.Val) + lll
		l1 = l1.Next
	}

	mmm := ""
  for l2 != nil {
		mmm = strconv.Itoa(l2.Val) + mmm
		l2 = l2.Next
  }

	// Convert strings to big integers
	llll := big.NewInt(0)
	llll.SetString(lll, 10)
	mmmm := big.NewInt(0)
	mmmm.SetString(mmm, 10)

	val := big.NewInt(0)
	val.Add(llll, mmmm)
	fmt.Println(val)

	aa := []ListNode{}

	//Convert to string, iterate thru nums
	valStr := val.String()
	fmt.Println(valStr[len(valStr)-1] - '0')

	// First value
	s := ListNode {
      Val: int(valStr[0] - '0'),
      Next: nil,
  }
  aa = append(aa, s)

  // all other values
	for j := 1; j < len(valStr); j++ {
		t := ListNode {
			Val: int(valStr[j] - '0'),
			Next: &aa[len(aa) - 1],
		}
		aa = append(aa, t)
	}

	return &aa[len(aa) - 1]
}
