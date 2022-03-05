package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	// Dummy node acts as the prevNode for the head node
	dummy := &ListNode{}
	dummy.Next = head
	prevNode := dummy

	for head != nil && head.Next != nil {

		// nodes that will be swapped
		firstNode := head
		secondNode := head.Next

		// Swap the two nodes. The swap step is
		firstNode.Next = secondNode.Next
		secondNode.Next = firstNode

		// We also need to assign the prevNode's next to the head of the swapped pair.ing
		prevNode.Next = secondNode

		// Reinitializing the head and prev
		prevNode = firstNode
		head = firstNode.Next

	}

	// Return the new head node.
	return dummy.Next
}
