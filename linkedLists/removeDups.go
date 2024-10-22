package linkedLists

import . "cracking-the-coding-interview/commons"

// Write code to remove duplicates from an unsorted linked list

func removeDuplicates(linkedList *LinkedList) {
	var duplicateChecker = make(map[int]int)

	actual := linkedList.GetHead()
	var previous *Node = nil

	for actual != nil {
		duplicateChecker[actual.Data] += 1

		if duplicateChecker[actual.Data] > 1 {
			previous.SetNext(actual.GetNext())

		}

		actual = actual.GetNext()
	}

}

func ExecuteRemoveDups() {

	node := NewNode(1)
	linkedList := NewLinkedList(*node)

	removeDuplicates(linkedList)

}
