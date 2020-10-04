package main

import (
	list "github.com/twsiyuan/study/cracking-the-code-review/chapter-2-linked-list/doubly"
)

// Palindrome:
// Implement a function to check if a linked list is a palindrome.
func IsPalindrome(l *list.List) bool {
	if l.Count() <= 0 {
		return false
	}

	a := l.First()
	b := l.Last()

	for a != nil && b != nil && a != b {
		if a.Value != b.Value {
			return false
		}
		a = a.Next()
		b = b.Previous()
	}
	return true
}
