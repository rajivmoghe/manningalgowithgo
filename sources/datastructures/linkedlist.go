package main

import (
	"fmt"
	"strings"
)

// Build a small linked list one cell at a time.
func small_list_test() {
	a_cell := Cell{data: "Apple"}
	b_cell := Cell{data: "Banana"}
	a_cell.next = &b_cell
	top := &a_cell

	for cell := top; cell != nil; cell = cell.next {
		fmt.Printf("%s ", cell.data)
	}
	fmt.Println()
}

// Add a cell after me.
func (me *Cell) add_after(after *Cell) {
	after.next = me.next
	me.next = after
}

// Delete the cell after me.
// Return the deleted cell.
func (me *Cell) delete_after() *Cell {
	if me.next != nil {
		delCell := me.next
		me.next = me.next.next
		return delCell
	}
	panic("no cell after me")
}

// Make a new LinkedList and initialize its sentinel.
func make_linked_list() LinkedList {
	list := LinkedList{}
	list.sentinel = &Cell{data: "SENTINEL"}
	return list
}

// Make a linked list containing the values.
func (list *LinkedList) add_range(values []string) {
	// get the last cell of the list
	lastCell := list.sentinel

	for {
		if lastCell.next != nil { // there is a cell after this one
			lastCell = lastCell.next
		} else { // no further cells, this is last. go out.
			break
		}
	}

	for _, value := range values { // loop over the slice
		newCell := &Cell{data: value}
		lastCell.add_after(newCell) // add thenew cell
		lastCell = newCell          // set last one 1 step forward
	}

}

// Return a string holding the cell values
// separated by the separator.
func (list *LinkedList) to_string(separator string) string {
	outStr := strings.Builder{}
	startCell := list.sentinel

	for startCell != nil {
		outStr.WriteString(startCell.data)
		outStr.WriteString(separator)
		startCell = startCell.next
	}

	return outStr.String()
}

// Return the number of cells in the list.
func (list *LinkedList) length() int {
	cellCount := 0
	startCell := list.sentinel

	for startCell != nil {
		cellCount++
		startCell = startCell.next
	}

	return cellCount
}
