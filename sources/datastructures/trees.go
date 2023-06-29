package main

import "strings"

// build a fixed tree.
func build_tree() *Node {
	a := Node{"A", nil, nil}
	b := Node{"B", nil, nil}
	c := Node{"C", nil, nil}
	d := Node{"D", nil, nil}
	e := Node{"E", nil, nil}
	f := Node{"F", nil, nil}
	g := Node{"G", nil, nil}
	h := Node{"H", nil, nil}
	i := Node{"I", nil, nil}
	j := Node{"J", nil, nil}

	a.left = &b
	b.left = &d
	b.right = &e
	e.left = &g
	a.right = &c
	c.right = &f
	f.left = &h
	h.left = &i
	h.right = &j

	return &a
}

// Display a tree with proper indentations
func (node *Node) display_indented(indent string, depth int) string {
	result := strings.Repeat(indent, depth) + node.node_data + "\n"
	if node.left != nil {
		result += node.left.display_indented(indent, depth+1)
	}
	if node.right != nil {
		result += node.right.display_indented(indent, depth+1)
	}

	return result
}

// Display a tree in preorder traversal
func (node *Node) preorder() string {
	result := " " + node.node_data
	if node.left != nil {
		result += node.left.preorder()
	}
	if node.right != nil {
		result += node.right.preorder()
	}

	return result
}

// Display a tree in inorder traversal
func (node *Node) inorder() string {
	result := ""
	if node.left != nil {
		result += node.left.inorder()
	}
	result += " " + node.node_data
	if node.right != nil {
		result += node.right.inorder()
	}

	return result
}

// Display a tree in postorder traversal
func (node *Node) postorder() string {
	result := ""
	if node.left != nil {
		result += node.left.postorder()
	}
	if node.right != nil {
		result += node.right.postorder()
	}
	result += " " + node.node_data

	return result
}

// Breadth-first traversals
func (node *Node) breadth_first() string {

	result := ""
	queue := make_queue_for_tree()
	queue.enqueue(node)

	// now dequeue the node, examine its children, and enqueue them
	for queue.top_sentinel.next != queue.btm_sentinel {
		procNode := queue.dequeue()
		result += procNode.node_data + " "
		if procNode.left != nil {
			queue.enqueue(procNode.left)
		}
		if procNode.right != nil {
			queue.enqueue(procNode.right)
		}
	}

	return result
}
