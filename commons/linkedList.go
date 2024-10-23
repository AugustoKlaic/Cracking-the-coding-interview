package commons

type LinkedList struct {
	head *Node
}

func NewLinkedList(nodes ...Node) *LinkedList {
	if len(nodes) == 0 {
		return &LinkedList{head: nil}
	}

	linkedList := &LinkedList{
		head: &nodes[0],
	}

	current := linkedList.head

	for i := 1; i < len(nodes); i++ {
		current.next = &nodes[i]
		current = current.next
	}

	return linkedList
}

func (list *LinkedList) GetHead() *Node {
	return list.head
}

func (list *LinkedList) appendToTail(data int) {
	end := NewNode(data)
	node := list.head

	for node.next != nil {
		node = node.next
	}

	node.next = end
}

func (list *LinkedList) deleteNode(head *Node, data int) *Node {
	if head == nil {
		return nil
	}

	node := head

	if node.Data == data {
		return head.next
	}

	for node.next != nil {
		if node.next.Data == data {
			node.next = node.next.next
			return head
		}
		node = node.next
	}
	return head
}
