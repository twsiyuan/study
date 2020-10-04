package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	list "github.com/twsiyuan/study/cracking-the-code-review/chapter-2-linked-list/doubly"
)

func IntArrayToList(ns ...int) *list.List {
	l := &list.List{}
	for _, n := range ns {
		l.PushLast(n)
	}
	return l
}

func TestConvertToN(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(123000, convertToN(IntArrayToList(0, 0, 0, 3, 2, 1)))
	assert.Equal(321, convertToN(IntArrayToList(1, 2, 3, 0, 0, 0)))
	assert.Equal(321, convertToN(IntArrayToList(1, 2, 3)))
	assert.Equal(1, convertToN(IntArrayToList(1)))

	assert.Equal(0, convertToN(StringToList("abc")))
	assert.Equal(0, convertToN(StringToList("123")))
	assert.Equal(0, convertToN(nil))
}

func TestConvertToNForward(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(321, convertToNForward(IntArrayToList(0, 0, 0, 3, 2, 1)))
	assert.Equal(123000, convertToNForward(IntArrayToList(1, 2, 3, 0, 0, 0)))
	assert.Equal(123, convertToNForward(IntArrayToList(1, 2, 3)))
	assert.Equal(1, convertToNForward(IntArrayToList(1)))
}
