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
