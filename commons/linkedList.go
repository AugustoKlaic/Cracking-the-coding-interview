package commons

type LinkedList struct {
	head *Node
}

func NewLinkedList(node Node) *LinkedList {
	return &LinkedList{
		head: &node,
	}
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

	if node.data == data {
		return head.next
	}

	for node.next != nil {
		if node.next.data == data {
			node.next = node.next.next
			return head
		}
		node = node.next
	}
	return head
}
