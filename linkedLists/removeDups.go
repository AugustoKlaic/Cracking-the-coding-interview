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
		} else {
			previous = actual
		}

		actual = actual.GetNext()
	}
}

func removeDuplicatesBookSolutionWithoutBuffer(linkedList *LinkedList) {
	actual := linkedList.GetHead()

	for actual != nil {
		runner := actual
		for runner.GetNext() != nil {
			if runner.GetNext().Data == actual.Data {
				runner.SetNext(runner.GetNext().GetNext())
			} else {
				runner = runner.GetNext()
			}
		}
		actual = actual.GetNext()
	}

}

func ExecuteRemoveDups() {

	node1 := NewNode(1)
	node2 := NewNode(1)
	node3 := NewNode(2)
	node4 := NewNode(3)

	linkedList := NewLinkedList(*node1, *node2, *node3, *node4)

	removeDuplicates(linkedList)

	node := linkedList.GetHead()

	for node != nil {
		println(node.Data)
		node = node.GetNext()
	}

}
