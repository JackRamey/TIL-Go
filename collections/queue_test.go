package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {

	t.Run("Add(value) increases the size of the queue", func(t *testing.T) {
		q := Queue[int]{}
		q.Add(1)
		assert.Len(t, q, 1)
	})

	t.Run("Add then Remove should return the original value pushed ", func(t *testing.T) {
		q := Queue[int]{}
		q.Add(8)
		assert.Equal(t, 8, q.Remove())
	})

	t.Run("Remove returns the first value pushed and removes it from the queue", func(t *testing.T) {
		q := Queue[int]{}
		q.Add(1)
		q.Add(2)
		assert.Len(t, q, 2)
		x := q.Remove()
		assert.Equal(t, x, 1)
		assert.Len(t, q, 1)
	})

	t.Run("Peek should return the the head of the queue but should not change the length", func(t *testing.T) {
		q := Queue[int]{}
		q.Add(7)
		l := len(q)
		assert.Equal(t, 7, q.Peek())
		assert.Equal(t, l, len(q))
	})

	t.Run("IsEmpty on a new queue returns true", func(t *testing.T) {
		q := Queue[int]{}
		assert.True(t, q.IsEmpty())
	})

	t.Run("IsEmpty on an empty queue returns true", func(t *testing.T) {
		q := Queue[int]{}
		q.Add(5)
		q.Remove()
		assert.True(t, q.IsEmpty())
	})

}
