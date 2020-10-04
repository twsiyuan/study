package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	list "github.com/twsiyuan/study/cracking-the-code-review/chapter-2-linked-list/doubly"
)

func StringToList(s string) *list.List {
	l := &list.List{}
	for _, r := range s {
		l.PushLast(r)
	}
	return l
}

func TestIsPalindrome(t *testing.T) {
	assert := assert.New(t)
	assert.True(IsPalindrome(StringToList("hannah")))
	assert.True(IsPalindrome(StringToList("otto")))
	assert.True(IsPalindrome(StringToList("bob")))
	assert.False(IsPalindrome(StringToList("")))
	assert.False(IsPalindrome(StringToList("siyuan")))
	assert.False(IsPalindrome(StringToList("london")))
}
