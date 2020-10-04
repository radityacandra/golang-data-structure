package datastructure

func NewStack() *Node {
	return &Node{}
}

func (node *Node) Push(data interface{}) {
	if node == nil {
		node.data = data
		return
	}

	newNode := Node{}
	newNode.data = data
	newNode.next = node

	*node = newNode
}

func (node *Node) Peek() interface{} {
	return node.data
}

func (node *Node) Pop() interface{} {
	data := node.data

	newNode := node.next
	node = newNode

	return data
}