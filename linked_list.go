package datastructure

type Node struct {
	next *Node
	data interface{}
}

func NewLinkedList() *Node {
	return &Node{}
}

func (node *Node) AddFirst(value interface{}) {
	if node.data == nil {
		node.data = value
	}

	newNode := Node{
		data: value,
	}

	*newNode.next = *node

	*node = newNode
}

func (node *Node) GetFirst() interface{} {
	return node.data
}

func (node *Node) GetLast() interface{} {
	current := *node
	for current.next != nil {
		current = *current.next
	}

	return current.data
}

func (node *Node) AddLast(value interface{}) {
	current := *node
	for current.next != nil {
		current = *current.next
	}

	newNode := Node{
		data: value,
	}

	*current.next = newNode
}

func (node *Node) Size() int {
	size := 0
	current := *node
	for current.next != nil {
		current = *current.next
		size++
	}

	return size
}

func (node *Node) Clear() {
	node = nil
}

func (node *Node) Delete(value interface{}) {
	previous := new(Node)
	current := *node
	for current.next != nil {
		if current.data == value {
			break
		}

		previous = &current
		current = *current.next
	}

	// if head
	if previous == nil {
		*node = *current.next
	}

	// if middle
	if previous != nil && current.next != nil {
		next := *current.next
		*previous.next = next

		*node = *previous
	}

	// if tail
	if current.next == nil {
		*node = *previous
	}
}