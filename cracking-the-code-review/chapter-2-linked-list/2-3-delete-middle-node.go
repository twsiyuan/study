package main

// Delete Middle Node:
//  Implement an algorithm to delete a node in the middle
// (i.e., any node but the first and last node, not necessarily the exact middle)
// of a singly linked list, given only access to that node.
func DeleteMiddleNode(head, node *Node) bool {
	if head == nil || node == nil {
		return false
	} else if head == node {
		panic("cannot delete first node")
	}

	at := head
	for ; at != nil; at = at.Next {
		if at.Next != node {
			continue
		}
		at.Next = node.Next // or at.Next.Next
		return true
	}
	return false
}
