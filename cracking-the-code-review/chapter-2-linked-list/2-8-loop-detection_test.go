package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetectLoop(t *testing.T) {
	assert := assert.New(t)

	t.Run("No loop", func(t *testing.T) {
		n := &Node{Value: "A"}
		n2 := MustAddNextNode(n, "B")
		n3 := MustAddNextNode(n2, "C")
		n4 := MustAddNextNode(n3, "D")
		_ = MustAddNextNode(n4, "E")

		assert.Nil(DetectLoop(n))
	})

	t.Run("Loop", func(t *testing.T) {
		n := &Node{Value: "A"}
		n2 := MustAddNextNode(n, "B")
		n3 := MustAddNextNode(n2, "C")
		n4 := MustAddNextNode(n3, "D")
		n5 := MustAddNextNode(n4, "E")
		n5.Next = n3

		assert.Equal(n3, DetectLoop(n))
	})
}
