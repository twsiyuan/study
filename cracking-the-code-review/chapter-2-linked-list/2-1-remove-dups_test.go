package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	list "github.com/twsiyuan/study/cracking-the-code-review/chapter-2-linked-list/singly"
)

func TestRemoveDups(t *testing.T) {
	buildList := func() *list.List {
		l := &list.List{}
		l.PushLast("alpha")
		l.PushLast("beta")
		l.PushLast("alpha")
		l.PushLast("beta")
		l.PushLast("charlie")
		l.PushLast("delta")
		return l
	}

	expected := []interface{}{
		"alpha",
		"beta",
		"charlie",
		"delta",
	}

	t.Run("With buffer", func(t *testing.T) {
		l := buildList()
		RemoveDups(l)

		assert.Equal(t, expected, list.ToArray(l))
	})

	t.Run("Without buffer", func(t *testing.T) {
		l := buildList()
		RemoveDupsNoBuffer(l)

		assert.Equal(t, expected, list.ToArray(l))
	})
}

func BenchmarkRemoveDups(b *testing.B) {
	buildList := func() *list.List {
		l := &list.List{}
		l.PushLast("alpha")
		l.PushLast("beta")
		l.PushLast("alpha")
		l.PushLast("beta")
		l.PushLast("charlie")
		l.PushLast("delta")
		return l
	}

	b.Run("With buffer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			l := buildList()
			RemoveDups(l)
		}
	})

	b.Run("Without buffer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			l := buildList()
			RemoveDupsNoBuffer(l)
		}
	})
}
