package linkedLists

import . "cracking-the-coding-interview/commons"

// Implement an algorithm to find the kth to last element of a singly linked list

func kthToLast(kthElement int, linkedList *LinkedList) *Node {
	var node = linkedList.GetHead()
	var count = 0
	var kthNode *Node = nil

	for node != nil {
		if count == kthElement {
			kthNode = linkedList.GetHead()
		}

		if kthNode != nil {
			kthNode = kthNode.GetNext()
		}

		node = node.GetNext()

		count++
	}

	if count == kthElement && kthNode == nil {
		return linkedList.GetHead()
	}

	return kthNode
}

func ExecuteKthToLast() {

	node1 := NewNode(1)
	node2 := NewNode(2)
	node3 := NewNode(3)
	node4 := NewNode(4)
	node5 := NewNode(5)
	node6 := NewNode(6)

	linkedList := NewLinkedList(*node1, *node2, *node3, *node4, *node5, *node6)

	node := kthToLast(1, linkedList)

	println("Kth element is: ", node.Data)

}
