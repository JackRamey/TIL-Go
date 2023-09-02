package collections

type Stack[T any] []T

func (stack *Stack[T]) IsEmpty() bool {
	return len(*stack) == 0
}

func (stack *Stack[T]) Peek() T {
	return (*stack)[len(*stack)-1]
}

func (stack *Stack[T]) Pop() T {
	hand := *stack
	var element T
	n := len(hand) - 1
	element, *stack = hand[n], hand[:n]
	return element
}

func (stack *Stack[T]) Push(value T) {
	*stack = append(*stack, value)
}
