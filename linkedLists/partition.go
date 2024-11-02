package linkedLists

import . "cracking-the-coding-interview/commons"

// Write code to partition a linked list around a value X, such that all nodes less than X come before all nodes greater
// than or equal to X. The partitionEasyMode element X can appear in any position (last, first, middle) in the right partitionEasyMode

func partition(list *LinkedList, partitionValue int) *LinkedList {
	var node, lastLeftNode *Node = list.GetHead(), nil
	var leftList, rightList = NewLinkedList(), NewLinkedList()

	for node != nil {
		if node.Data < partitionValue {
			auxNode := NewNode(node.Data)
			leftList.AppendToTail(auxNode)
			lastLeftNode = auxNode
		} else {
			rightList.AppendToTail(NewNode(node.Data))
		}
		node = node.GetNext()
	}

	lastLeftNode.SetNext(rightList.GetHead())
	return leftList
}

func ExecutePartition() {

	node1 := NewNode(2)
	node2 := NewNode(7)
	node3 := NewNode(3)
	node4 := NewNode(8)
	node5 := NewNode(5)
	node6 := NewNode(1)

	linkedList := NewLinkedList(*node1, *node2, *node3, *node4, *node5, *node6)

	leftList := partition(linkedList, 4)

	leftList.PrintList()
}
