package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	list "github.com/twsiyuan/study/cracking-the-code-review/chapter-2-linked-list/doubly"
)

func TestPartition(t *testing.T) {
	assert := assert.New(t)

	t.Run("Case 1", func(t *testing.T) {
		input := IntArrayToList(3, 5, 8, 5, 10, 2, 1)
		output := Partition(input, 11)

		assert.Equal([]interface{}{3, 5, 8, 5, 10, 2, 1}, list.ToArray(output))
	})

	t.Run("Case 2", func(t *testing.T) {
		input := IntArrayToList(3, 5, 8, 5, 10, 2, 1)
		output := Partition(input, 0)

		assert.Equal([]interface{}{3, 5, 8, 5, 10, 2, 1}, list.ToArray(output))
	})

	t.Run("Case 3", func(t *testing.T) {
		input := IntArrayToList(3, 5, 8, 5, 10, 2, 1)
		output := Partition(input, 5)

		assert.Equal([]interface{}{3, 2, 1, 5, 8, 5, 10}, list.ToArray(output))
	})
}
