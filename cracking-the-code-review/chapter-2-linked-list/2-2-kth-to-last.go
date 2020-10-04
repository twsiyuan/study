package main

import (
	list "github.com/twsiyuan/study/cracking-the-code-review/chapter-2-linked-list/singly"
)

// Return Kth to Last:
// Implement an algorithm to find the kth to last element of a singly linked list.
func KthToLast(l *list.List, k int) []interface{} {
	c := l.Count() - k
	if c <= 0 {
		return []interface{}{}
	}

	at := l.First()
	for i := 0; i < k && at != nil; i++ {
		at = at.Next()
	}

	result := make([]interface{}, 0, c)
	for ; at != nil; at = at.Next() {
		result = append(result, at.Value)
	}
	return result
}
