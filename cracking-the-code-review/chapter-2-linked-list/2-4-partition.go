package main

import (
	list "github.com/twsiyuan/study/cracking-the-code-review/chapter-2-linked-list/doubly"
)

// Partition:
// Write code to partition a linked list around a value x,
// such that all nodes less than x come before all nodes greater than or equal to x.
// lf x is contained within the list,
// the values of x only need to be after the elements less than x (see below).
// The partition element x can appear anywhere in the "right partition";
// it does not need to appear between the left and right partitions.
//
// INPUT:  3 -> 5 -> 8 -> 5 ->10 -> 2 -> 1[partition=5)
// OUTPUT: 3 -> 1 -> 2 -> 10 -> 5 -> 5 -> 8
//
func Partition(l *list.List, num int) *list.List {
	if l == nil {
		return nil
	}

	right := &list.List{}
	left := &list.List{}
	for at := l.First(); at != nil; at = at.Next() {
		n, ok := at.Value.(int)
		if !ok {
			panic("cannot partition non-integer list")
		}
		if n < num {
			left.PushLast(n)
		} else {
			right.PushLast(n)
		}
	}

	left.PushLastList(right)
	return left
}
