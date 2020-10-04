package main

// Intersection:
// Given two (singly) linked lists,
// determine if the two lists intersect.
// Return the intersecting node.
// Note that the intersection is defined based on reference, not value.
//That is, if the kth node of the first linked list is the exact same node
// (by reference) as the jth node of the second linked list, then they are intersecting.
func IsIntersecting(listA, listB *Node) *Node {
	for atK := listA; atK != nil; atK = atK.Next {
		for atJ := listB; atJ != nil; atJ = atJ.Next {
			if atK == atJ {
				return atK
			}
		}
	}
	return nil
}
