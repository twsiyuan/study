package main

import (
	list "github.com/twsiyuan/study/cracking-the-code-review/chapter-2-linked-list/singly"
)

// Remove Dups:
// Write code to remove duplicates from an unsorted linked list.
func RemoveDups(l *list.List) {
	dups := make(map[interface{}]int, l.Count())
	for e := l.First(); e != nil; {
		next := e.Next()
		if _, ok := dups[e.Value]; ok {
			l.Remove(e)
			e = next
			continue
		}

		dups[e.Value] = 1
		e = next
	}
}

// Remove Dups:
// Write code to remove duplicates from an unsorted linked list. FOLLOW UP
// How would you solve this problem if a temporary buffer is not allowed?
func RemoveDupsNoBuffer(l *list.List) {
	for e := l.First(); e != nil; e = e.Next() {
		for a := e.Next(); a != nil; {
			next := a.Next()
			if a.Value == e.Value {
				l.Remove(a)
			}
			a = next
		}
	}
}
