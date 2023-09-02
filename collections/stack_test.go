package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {

	t.Run("Push(value) increases the size of the stack", func(t *testing.T) {
		s := Stack[int]{}
		s.Push(1)
		assert.Len(t, s, 1)
	})

	t.Run("Push then Pop should return the original value pushed ", func(t *testing.T) {
		s := Stack[int]{}
		s.Push(8)
		assert.Equal(t, 8, s.Pop())
	})

	t.Run("Pop returns the last value pushed and removes it from the stack", func(t *testing.T) {
		s := Stack[int]{}
		s.Push(1)
		s.Push(2)
		assert.Len(t, s, 2)
		x := s.Pop()
		assert.Equal(t, x, 2)
		assert.Len(t, s, 1)
	})

	t.Run("Peek should return the the head of the stack but should not change the length", func(t *testing.T) {
		s := Stack[int]{}
		s.Push(7)
		l := len(s)
		assert.Equal(t, 7, s.Peek())
		assert.Equal(t, l, len(s))
	})

	t.Run("IsEmpty on a new stack returns true", func(t *testing.T) {
		s := Stack[int]{}
		assert.True(t, s.IsEmpty())
	})

	t.Run("IsEmpty on an empty stack returns true", func(t *testing.T) {
		s := Stack[int]{}
		s.Push(5)
		s.Pop()
		assert.True(t, s.IsEmpty())
	})

}
