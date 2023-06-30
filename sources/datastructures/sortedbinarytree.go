package main

func (node *Node) insert_value(new_data string) {
	if new_data == "" {
		panic("Cannot add a nil to the tree")
	}
	for {
		if new_data <= node.node_data {
			if node.left == nil {
				node.left = &Node{node_data: new_data}
				return
			} else {
				node = node.left
			}
		} else {
			if node.right == nil {
				node.right = &Node{node_data: new_data}
				return
			} else {
				node = node.right
			}
		}
	}

}

func (node *Node) find_value(new_data string) *Node {
	if node.node_data == new_data {
		return node
	}
	if node.node_data > new_data {
		if node.left != nil {
			return node.left.find_value(new_data)
		}
		return nil
	}
	if node.node_data < new_data {
		if node.right != nil {
			return node.right.find_value(new_data)
		}
		return nil
	}
	return nil
}
