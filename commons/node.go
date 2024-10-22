package commons

type Node struct {
	next *Node
	Data int
}

func (node *Node) GetNext() *Node {
	return node.next
}

func (node *Node) SetNext(next *Node) {
	node.next = next
}

func NewNode(data int) *Node {
	return &Node{
		next: nil,
		Data: data,
	}
}
