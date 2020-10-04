package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsIntersecting(t *testing.T) {
	assert := assert.New(t)
	t.Run("Nil", func(t *testing.T) {
		n := &Node{}
		assert.Nil(IsIntersecting(nil, nil))
		assert.Nil(IsIntersecting(nil, n))
		assert.Nil(IsIntersecting(n, nil))
	})

	t.Run("No intersection", func(t *testing.T) {
		n := &Node{Value: "a"}
		n2 := MustAddNextNode(n, "b")
		n3 := MustAddNextNode(n2, "c")
		_ = MustAddNextNode(n3, "d")

		m := &Node{Value: "a"}
		m2 := MustAddNextNode(m, "b")
		m3 := MustAddNextNode(m2, "c")
		_ = MustAddNextNode(m3, "d")

		assert.Nil(IsIntersecting(n, m))
	})

	t.Run("Intersection", func(t *testing.T) {
		n := &Node{Value: "a"}
		n2 := MustAddNextNode(n, "b")
		n3 := MustAddNextNode(n2, "c")
		n4 := MustAddNextNode(n3, "d")

		m := &Node{Value: "a"}
		m2 := MustAddNextNode(m, "b")
		m2.Next = n2

		assert.Equal(n2, IsIntersecting(n, m))
		assert.Equal(n4, IsIntersecting(n, n4))
	})
}
