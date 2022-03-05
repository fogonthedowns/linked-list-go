package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	ptr := head
	var ktail *ListNode

	// head of the final, modified linked list
	var new_head *ListNode

	// loop through all the nodes
	for ptr != nil {
		count := 0

		// start counting from the head
		ptr = head

		// find the head of the next k nodes
		for count < k && ptr != nil {
			ptr = ptr.Next
			count += 1
		}

		// if we counted up to k, reverse them!
		if count == k {

			// Reverse k nodes and get the new head
			revHead := reverseLinkedList(head, k)

			// new_head is the head of the final linked list
			if new_head == nil {
				new_head = revHead
			}

			// ktail is the tail of the previous block of
			// revered k nodes
			if ktail != nil {
				ktail.Next = revHead
			}

			// set ktail for next iteration
			ktail = head
			// set head for next iteration
			head = ptr
		}
	}

	// attach the final, possibly un-reversed portion
	if ktail != nil {
		ktail.Next = head
	}

	if new_head == nil {
		return head
	} else {
		return new_head
	}

}

func reverseLinkedList(head *ListNode, k int) *ListNode {
	var new_head *ListNode
	ptr := head

	for k > 0 {
		// keep track of the next node
		next_node := ptr.Next

		ptr.Next = new_head
		new_head = ptr

		ptr = next_node
		k--
	}
	return new_head
}
