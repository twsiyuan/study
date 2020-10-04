package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	list "github.com/twsiyuan/study/cracking-the-code-review/chapter-2-linked-list/singly"
)

func TestKthToLast(t *testing.T) {
	l := &list.List{}
	l.PushLast(0)
	l.PushLast(1)
	l.PushLast(2)
	l.PushLast(3)
	l.PushLast(4)

	assert := assert.New(t)
	assert.Equal([]interface{}{0, 1, 2, 3, 4}, KthToLast(l, -1))
	assert.Equal([]interface{}{0, 1, 2, 3, 4}, KthToLast(l, 0))
	assert.Equal([]interface{}{2, 3, 4}, KthToLast(l, 2))
	assert.Equal([]interface{}{3, 4}, KthToLast(l, 3))
	assert.Equal([]interface{}{}, KthToLast(l, 6))
}
