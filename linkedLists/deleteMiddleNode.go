package linkedLists

import . "cracking-the-coding-interview/commons"

// Implement an algorithm to delete a node in the middle (all the nodes between the first and last) of singly linked list,
// given only access to that node

func deleteMiddleNode(linkedList *LinkedList, indexToDelete int) {
	var node = linkedList.GetHead()
	var previousNode *Node = nil

	var count = 1

	for node != nil {
		if count == indexToDelete {
			if node.GetNext() != nil && count > 1 {
				previousNode.SetNext(node.GetNext())
			}
		} else {
			previousNode = node
		}

		node = node.GetNext()
		count++
	}
}

func deleteMiddleNodeOnlyAccessingTheNode(node *Node) {
	if node.GetNext() != nil || node != nil {
		var next *Node = node.GetNext()
		node.Data = next.Data
		node.SetNext(next.GetNext())
	}
}

func ExecuteDeleteMiddleNode() {
	node1 := NewNode(1)
	node2 := NewNode(2)
	node3 := NewNode(3)
	node4 := NewNode(4)
	node5 := NewNode(5)
	node6 := NewNode(6)

	linkedList := NewLinkedList(*node1, *node2, *node3, *node4, *node5, *node6)

	deleteMiddleNode(linkedList, 2)

	PrintList(linkedList)

}
