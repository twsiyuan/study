package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func ToArrayFromFirst(l *List) []interface{} {
	return ToArray(l)
}

func ToArrayFromLast(l *List) []interface{} {
	a := make([]interface{}, 0, l.Count())
	for at := l.Last(); at != nil; at = at.Previous() {
		a = append(a, at.Value)
	}
	return a
}

func TestPushLast(t *testing.T) {
	assert := assert.New(t)

	t.Run("Nil", func(t *testing.T) {
		l := (*List)(nil)
		assert.Nil(l.PushLast(1))
	})

	t.Run("OK", func(t *testing.T) {
		l := &List{}
		assert.Equal(0, l.Count())
		assert.Nil(l.First())
		assert.Nil(l.Last())
		assert.Equal([]interface{}{}, ToArrayFromFirst(l))
		assert.Equal([]interface{}{}, ToArrayFromLast(l))

		l.PushLast(1)
		assert.Equal(1, l.Count())
		assert.Equal(1, l.First().Value)
		assert.Equal(1, l.Last().Value)
		assert.Equal([]interface{}{1}, ToArrayFromFirst(l))
		assert.Equal([]interface{}{1}, ToArrayFromLast(l))

		l.PushLast(2)
		assert.Equal(2, l.Count())
		assert.Equal(1, l.First().Value)
		assert.Equal(2, l.Last().Value)
		assert.Equal([]interface{}{1, 2}, ToArrayFromFirst(l))
		assert.Equal([]interface{}{2, 1}, ToArrayFromLast(l))

		l.PushLast(3)
		assert.Equal(3, l.Count())
		assert.Equal(1, l.First().Value)
		assert.Equal(3, l.Last().Value)
		assert.Equal([]interface{}{1, 2, 3}, ToArrayFromFirst(l))
		assert.Equal([]interface{}{3, 2, 1}, ToArrayFromLast(l))
	})
}

func TestPushLastSlice(t *testing.T) {
	assert := assert.New(t)

	t.Run("OK", func(t *testing.T) {
		l := &List{}
		l.PushLastSlice(1, 2, 3, 4, 5)
		assert.Equal([]interface{}{1, 2, 3, 4, 5}, ToArrayFromFirst(l))
		assert.Equal([]interface{}{5, 4, 3, 2, 1}, ToArrayFromLast(l))
	})
}

func TestRemove(t *testing.T) {
	assert := assert.New(t)

	t.Run("Edge cases", func(t *testing.T) {
		l := &List{}
		assert.NotPanics(func() {
			assert.False(l.Remove(nil))
			assert.False(l.Remove(&Element{}))
		})

		l = nil
		assert.NotPanics(func() {
			assert.False(l.Remove(&Element{}))
		})
	})

	t.Run("Element does not exist", func(t *testing.T) {
		l := &List{}
		l.PushLast(1)
		l.PushLast(2)
		l.PushLast(3)

		assert.False(l.Remove(nil))
		assert.False(l.Remove(&Element{}))
	})

	t.Run("OK", func(t *testing.T) {
		l := &List{}
		e1 := l.PushLast(1)
		e2 := l.PushLast(2)
		e3 := l.PushLast(3)

		assert.True(l.Remove(e3))
		assert.Equal(2, l.Count())
		assert.Equal(e1, l.First())
		assert.Equal(e2, l.Last())
		assert.Equal([]interface{}{1, 2}, ToArrayFromFirst(l))
		assert.Equal([]interface{}{2, 1}, ToArrayFromLast(l))

		assert.True(l.Remove(e2))
		assert.Equal(1, l.Count())
		assert.Equal(e1, l.First())
		assert.Equal(e1, l.Last())
		assert.Equal([]interface{}{1}, ToArrayFromFirst(l))
		assert.Equal([]interface{}{1}, ToArrayFromLast(l))

		assert.False(l.Remove(e2))

		assert.True(l.Remove(e1))
		assert.Equal(0, l.Count())
		assert.Nil(l.First())
		assert.Nil(l.Last())
		assert.Equal([]interface{}{}, ToArrayFromFirst(l))
		assert.Equal([]interface{}{}, ToArrayFromLast(l))
	})
}
