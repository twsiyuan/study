package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteMiddleNode(t *testing.T) {
	assert := assert.New(t)

	head := &Node{Value: "a"}
	nodeB := MustAddNextNode(head, "b")
	nodeC := MustAddNextNode(nodeB, "c")
	nodeD := MustAddNextNode(nodeC, "d")
	nodeE := MustAddNextNode(nodeD, "e")
	nodeF := MustAddNextNode(nodeE, "f")

	assert.True(DeleteMiddleNode(head, nodeC))
	assert.Equal([]interface{}{"a", "b", "d", "e", "f"}, ToArray(head))
	assert.False(DeleteMiddleNode(head, nodeC))

	assert.True(DeleteMiddleNode(head, nodeF))
	assert.Equal([]interface{}{"a", "b", "d", "e"}, ToArray(head))

	assert.True(DeleteMiddleNode(head, nodeB))
	assert.Equal([]interface{}{"a", "d", "e"}, ToArray(head))
}

func MustAddNextNode(h *Node, value interface{}) *Node {
	if h == nil {
		panic("h is nil")
	} else if h.Next != nil {
		panic("next does exist")
	}
	m := &Node{Value: value}
	h.Next = m
	return m
}
