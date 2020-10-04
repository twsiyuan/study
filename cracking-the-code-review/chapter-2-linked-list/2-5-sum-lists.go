package main

import (
	"math"

	list "github.com/twsiyuan/study/cracking-the-code-review/chapter-2-linked-list/doubly"
)

// Sum Lists:
// You have two numbers represented by a linked list,
// where each node contains a single digit.
// The digits are stored in reverse order,
// such that the 1's digit is at the head of the list.
// Write a function that adds the two numbers and returns the sum as a linked list.
//
// (7-> 1 -> 6) + (5 -> 9 -> 2) = 617 + 592
//
func SumLists(a, b *list.List) int {
	return convertToN(a) + convertToN(b)
}

func convertToN(l *list.List) int {
	if l == nil {
		return 0
	}

	n := 0
	digit := 0
	for at := l.First(); at != nil; at = at.Next() {
		m, ok := at.Value.(int)
		if !ok {
			return 0
		}
		n += (int)(math.Pow10(digit)) * m
		digit++
	}
	return n
}

func convertToNForward(l *list.List) int {
	if l == nil {
		return 0
	}

	n := 0
	digit := l.Count() - 1
	for at := l.First(); at != nil; at = at.Next() {
		m, ok := at.Value.(int)
		if !ok {
			return 0
		}
		n += (int)(math.Pow10(digit)) * m
		digit--
	}
	return n
}
