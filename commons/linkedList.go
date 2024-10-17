package commons

type LinkedList struct {
	head *Node
}

func NewLinkedList(node Node) *LinkedList {
	return &LinkedList{
		head: &node,
	}
}
