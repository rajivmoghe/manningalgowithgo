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

type Node struct {
	node_data   string
	left, right *Node
}

type TreeCell struct {
	tree_node *Node
	next      *TreeCell
	prev      *TreeCell
}

type TreeQueue struct {
	top_sentinel *TreeCell
	btm_sentinel *TreeCell
}

