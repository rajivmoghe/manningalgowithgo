package main

import "strings"

// Return true if the list has a loop.
func (list *LinkedList) has_loop() bool {
	// start at the sentinel
	slow := list.sentinel
	fast := list.sentinel

	for {
		if fast == nil || fast.next == nil { // no more links left
			return false
		}

		fast = fast.next.next
		slow = slow.next

		if fast == slow { // pointers have met again
			return true
		}
	}

}

func (list *LinkedList) to_string_max(separator string, max int) string {
	outStr := strings.Builder{}
	// outStr.WriteString(list.sentinel.data)
	node := list.sentinel.next

	for i := 0; i < max; i++ {
		if node == nil {
			break
		}
		outStr.WriteString(node.data)
		outStr.WriteString(separator)
		node = node.next
	}

	return outStr.String()
}
