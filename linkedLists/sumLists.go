package linkedLists

import . "cracking-the-coding-interview/commons"

// Two numbers represented by a linked list, where each node contains a single digit. The digits are stored in reverse order,
// such that the first digit is at the head of the list. Write a function that adds the two numbers and returns the sum
// as a linked list (do not convert the list to an integer)
// EX: [7] [1] [6] + [5] [9] [2] = [2] [1] [9] -> 617 + 295 = 912

// Follow up: suppose the digits are stored in forward order. Repeat the above problem.
// EX: [6] [1] [7] + [2] [9] [5] = [9] [1] [2] -> 617 + 295 = 912

func sumListsFollowUp(firstList, secondList *LinkedList) *LinkedList {
	var lenght1, lenght2 = firstList.Size(), secondList.Size()

	if lenght1 > lenght2 {
		PadList(secondList, lenght1-lenght2)
	}

	if lenght2 > lenght1 {
		PadList(firstList, lenght2-lenght1)
	}

	var carry, resultList = sumHelper(firstList.GetHead(), secondList.GetHead())

	if carry > 0 {
		resultList.AppendToTail(NewNode(1))
		return resultList
	} else {
		return resultList
	}
}

func sumHelper(firstNode, secondNode *Node) (carry int, result *LinkedList) {
	if firstNode == nil && secondNode == nil {
		return 0, NewLinkedList()
	}

	nextCarry, nextResult := sumHelper(firstNode.GetNext(), secondNode.GetNext())

	sum := firstNode.Data + secondNode.Data + nextCarry

	if sum > 9 {
		sum = sum - 10
		carry = 1
	} else {
		carry = 0
	}

	nextResult.AppendToTail(NewNode(sum))

	return carry, nextResult
}

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

	if overflow {
		result.AppendToTail(NewNode(1))
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
	PrintList(sumListsFollowUp(firstList, secondList))
}
