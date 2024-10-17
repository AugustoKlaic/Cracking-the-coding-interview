package commons

type Node struct {
	next *Node
	data int
}

func NewNode(data int) *Node {
	return &Node{
		next: nil,
		data: data,
	}
}

func (actual *Node) appendToTail(data int) {
	end := NewNode(data)

	for actual.next != nil {
		actual = actual.next
	}

	actual.next = end
}
