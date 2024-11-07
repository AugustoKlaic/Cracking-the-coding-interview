package linkedLists

import . "cracking-the-coding-interview/commons"

// Two numbers represented by a linked list, where each node contains a single digit. The digits are stored in reverse order,
// such that the first digit is at the head of the list. Write a function that adds the two numbers and returns the sum
// as a linked list (do not convert the list to an integer)
// EX: [7] [1] [6] + [5] [9] [2] = [2] [1] [9] -> 617 + 295 = 912

func sumLists(firstList, secondList *LinkedList) *LinkedList {
	var nodeToSumOne, nodeToSumTwo *Node = firstList.GetHead(), secondList.GetHead()
	var result = NewLinkedList()
	var overflow = false

	for nodeToSumOne != nil || nodeToSumTwo != nil {
		var valueToSumOne, valueToSumTwo = 0, 0
		var dataSum = 0

		if nodeToSumOne != nil {
			valueToSumOne = nodeToSumOne.Data
			nodeToSumOne = nodeToSumOne.GetNext()
		}

		if nodeToSumTwo != nil {
			valueToSumTwo = nodeToSumTwo.Data
			nodeToSumTwo = nodeToSumTwo.GetNext()
		}

		dataSum = valueToSumOne + valueToSumTwo

		if overflow {
			dataSum++
		}

		if dataSum > 9 {
			overflow = true
			dataSum = dataSum - 10
		} else {
			overflow = false
		}

		result.AppendToTail(NewNode(dataSum))
	}

	return result
}

func ExecuteSumLists() {
	node1 := NewNode(7)
	node2 := NewNode(1)
	node3 := NewNode(6)

	node4 := NewNode(5)
	node5 := NewNode(9)
	node6 := NewNode(2)

	firstList := NewLinkedList(*node1, *node2, *node3)
	secondList := NewLinkedList(*node4, *node5, *node6)

	PrintList(sumLists(firstList, secondList))
}
