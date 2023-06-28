package main

import (
	"fmt"
	"strings"
)

// smal list testfor doubly linked list
func small_dl_list_test() {
	a_cell := Cell{data: "Apple"}
	b_cell := Cell{data: "Banana"}
	a_cell.next = &b_cell
	b_cell.prev = &a_cell
	top := &a_cell

	for cell := top; cell != nil; cell = cell.next {
		fmt.Printf("%s \n\tprev--> %s \n\n\tnext--> %s\n", cell.data, &cell.prev, cell.next)
	}
	fmt.Println()
}

// Make a new DoublyLinkedList and initialize its sentinels
func make_doubly_linked_list() DoublyLinkedList {
	list := DoublyLinkedList{}
	top_sentinel := &Cell{data: "START"}
	btm_sentinel := &Cell{data: "END"}

	top_sentinel.prev = nil
	top_sentinel.next = btm_sentinel

	btm_sentinel.prev = top_sentinel
	btm_sentinel.next = nil

	list.top_sentinel = top_sentinel
	list.btm_sentinel = btm_sentinel

	return list
}

// Add a cell after me.
func (me *Cell) dll_add_after(after *Cell) {
	after.next = me.next
	me.next.prev = after

	after.prev = me
	me.next = after
}

// Add a cell before me.
func (me *Cell) dll_add_before(before *Cell) {
	me.prev.dll_add_after(before)
}

// Delete the cell and return it.
func (me *Cell) dll_delete() *Cell {
	if me == nil {
		panic("Cell is nil. Can't remove.")
	}
	me.prev.next = me.next
	me.next.prev = me.prev

	me.prev = nil
	me.next = nil

	return me
}

// Add a range of values to an existing DLL
func (dll *DoublyLinkedList) add_range(values []string) {
	// First - get the last Cell
	handle := dll.btm_sentinel

	for _, valu := range values {
		newCell := &Cell{data: valu}
		handle.dll_add_before(newCell)
	}
}

// Return a string holding the cell values
func (dll *DoublyLinkedList) to_string(separator string) string {
	outStr := strings.Builder{}
	handle := dll.top_sentinel.next

	for handle != dll.btm_sentinel {
		outStr.WriteString(handle.data)
		outStr.WriteString(separator)
		handle = handle.next
	}
	return outStr.String()
}

// Return the number of cells in the list.
func (list *DoublyLinkedList) length() int {
	counter := 0
	countingCell := list.top_sentinel.next
	for countingCell != list.btm_sentinel {
		counter++
		countingCell = countingCell.next
	}
	return counter
}

// Return true if the queue is empty, false otherwise.
func (queue *DoublyLinkedList) is_empty() bool {
	return queue.top_sentinel.next == queue.btm_sentinel
}

// *** Queue functions ***

// Push an item at the back of the queue right before btm_sentinel
func (queue *DoublyLinkedList) push(value string) {
	queue.add_range([]string{value})
}

// Pop an item off of the front of the queue right after the top_sentinel.
func (queue *DoublyLinkedList) pop() string {
	return queue.top_sentinel.delete_after().data
}

// Add an item to the ~top~ back of the queue.
func (queue *DoublyLinkedList) enqueue(value string) {
    queue.push(value)
}

// Remove an item from the ~bottom~ front of the queue.
func (queue *DoublyLinkedList) dequeue() string {
    if queue.is_empty() {
        panic("Cannot remove items from empty queue")
    }
	retCell := queue.top_sentinel.next.dll_delete()
	return retCell.data
}

// *** Deque functions ***

// Add an item at the ~top~ back of the deque.
func (deque *DoublyLinkedList) push_top(value string) {
    deque.enqueue(value)
}

// Add an item at the ~bottom~ front of the deque.
func (deque *DoublyLinkedList) push_bottom(value string) {
    deque.top_sentinel.dll_add_after(&Cell{data: value})
}

// Remove an item from the ~top~ back of the deque.
func (deque *DoublyLinkedList) pop_top() string {
    if deque.is_empty() {
        panic("Cannot remove items from empty queue")
    }
	retCell := deque.btm_sentinel.prev.dll_delete()
	return retCell.data
}

// Add an item at the ~bottom~ front of the deque.
func (deque *DoublyLinkedList) pop_bottom() string {
    if deque.is_empty() {
        panic("Cannot remove items from empty queue")
    }
	return deque.dequeue()
}

