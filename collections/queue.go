package collections

// Queue is a First-In-First-Out data structure
type Queue[T any] []T

func (queue *Queue[T]) IsEmpty() bool {
	return len(*queue) == 0
}

// Peek returns the item at the head of the queue without removing it from its position.
func (queue *Queue[T]) Peek() T {
	return (*queue)[0]
}

// Remove returns the item at the head of the queue and then removes it from the leading position.
func (queue *Queue[T]) Remove() T {
	hand := *queue
	var element T
	length := len(hand)
	element, *queue = hand[0], hand[1:length]
	return element
}

// Add adds the value to the end of the queue.
func (queue *Queue[T]) Add(value T) {
	*queue = append(*queue, value)
}
