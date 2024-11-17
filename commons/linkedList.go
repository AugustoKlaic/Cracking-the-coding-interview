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

func (list *LinkedList) AppendToTail(newNode *Node) {
	end := newNode

	if list.head == nil {
		list.head = newNode
	} else {
		node := list.head

		for node.next != nil {
			node = node.next
		}
		node.next = end
	}
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

func (list *LinkedList) Size() int {
	var size int = 0
	var node *Node = list.head

	for node != nil {
		size++
		node = node.GetNext()
	}

	return size
}

func PadList(list *LinkedList, padSize int) {

	for i := 0; i < padSize; i++ {
		var oldHead *Node = list.head
		var newHead = NewNode(0)
		newHead.SetNext(oldHead)

		list.head = newHead
	}
}
