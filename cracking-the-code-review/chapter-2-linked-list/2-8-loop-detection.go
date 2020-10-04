package main

// Loop Detection
// Given a circular linked list, implement an algorithm that
// returns the node at the beginning of the loop.
//
// DEFINITION
// Circular linked list: A (corrupt) linked list in which a node's next pointer points
// to an earlier node, so as to make a loop in the linked list.
func DetectLoop(n *Node) *Node {
	m := make(map[*Node]bool)
	for at := n; at != nil; at = at.Next {
		if _, ok := m[at]; ok {
			return at
		}
		m[at] = true
	}
	return nil
}
