package main

// *** Stack functions ***

// Push an item onto the top of the list right after the sentinel.
func (stack *LinkedList) push(value string) {
	stack.add_range([]string{value})
}

// Pop an item off of the list (from right after the sentinel).
func (stack *LinkedList) pop() string {
	return stack.sentinel.delete_after().data
}

// Return true if the stack is empty, false otherwise.
func (stack *LinkedList) is_empty() bool {
	return stack.sentinel.next == nil
}
