package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	length int
	start  *Node
	end    *Node
}

// Append/Prepend O(1)
func (ll *LinkedList) Push(n *Node) {
	if ll.length == 0 {
		ll.start = n
		ll.end = n
	} else {
		last := ll.end
		last.next = n
		ll.end = n
	}
	ll.length++
}

// Access/Search O(n)
func (ll *LinkedList) Pop(index int) *Node {
	if ll.length == 0 {
		return nil
	}

	var previous *Node
	current := ll.start

	for i := 0; i < index; i++ {
		if current.next == nil {
			return nil
		}

		previous = current
		current = current.next
	}
	previous.next = current.next
	ll.length--
	return current
}

// Delete O(N)
func (ll *LinkedList) Delete(index int) {
	if ll.length == 0 {
		return
	}

	var previous *Node
	current := ll.start

	for i := 0; i <= index; i++ {
		fmt.Printf("%v \n", current)
		if current.next == nil {
			return
		}

		previous = current
		current = current.next
	}
	previous.next = current.next
	ll.length--
}
