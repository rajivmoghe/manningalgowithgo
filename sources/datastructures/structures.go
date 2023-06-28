package main

type Cell struct {
	data string
	next *Cell
	prev *Cell
}

type LinkedList struct {
	sentinel *Cell
}

type DoublyLinkedList struct {
	top_sentinel *Cell
	btm_sentinel *Cell
}
