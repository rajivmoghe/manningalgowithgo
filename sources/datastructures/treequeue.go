package main

// Queue based on doubly linked lists for TreeNodes, instead of Cells.
// We do not need this if the doubly linked list used generics.

// Make a new DoublyLinkedList and initialize its sentinels
func make_queue_for_tree() TreeQueue {
	tree_queue := TreeQueue{}
	top_sentinel := &TreeCell{tree_node: &Node{"top_sentinel", nil, nil}}
	btm_sentinel := &TreeCell{tree_node: &Node{"top_sentinel", nil, nil}}

	top_sentinel.prev = nil
	top_sentinel.next = btm_sentinel

	btm_sentinel.prev = top_sentinel
	btm_sentinel.next = nil

	tree_queue.top_sentinel = top_sentinel
	tree_queue.btm_sentinel = btm_sentinel

	return tree_queue
}

// Add an item to the ~top~ back of the queue.
func (queue *TreeQueue) enqueue(node *Node) {

	handle := queue.btm_sentinel
	newTreeCell := &TreeCell{tree_node: node}
	handle.dll_add_before(newTreeCell)
}

// Remove an item from the ~bottom~ front of the queue.
func (queue *TreeQueue) dequeue() *Node {
	if queue.is_empty() {
		panic("Cannot remove items from empty queue")
	}
	retNode := queue.top_sentinel.next.dll_delete()
	return retNode.tree_node
}

// Return true if the queue is empty, false otherwise.
func (queue *TreeQueue) is_empty() bool {
	return queue.top_sentinel.next == queue.btm_sentinel
}

// Add a cell before me.
func (me *TreeCell) dll_add_before(before *TreeCell) {
	me.prev.dll_add_after(before)
}

// Add a cell after me.
func (me *TreeCell) dll_add_after(after *TreeCell) {
	after.next = me.next
	me.next.prev = after

	after.prev = me
	me.next = after
}

// Delete the cell and return it.
func (me *TreeCell) dll_delete() *TreeCell {
	if me == nil {
		panic("Cell is nil. Can't remove.")
	}
	me.prev.next = me.next
	me.next.prev = me.prev

	me.prev = nil
	me.next = nil

	return me
}
