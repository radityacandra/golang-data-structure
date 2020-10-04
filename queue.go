package datastructure

type Queue struct {
	tail *Node
	head *Node
}

func NewQueue() *Queue {
	return &Queue{}
}

func (queue *Queue) Add(value interface{}) {
	newNode := Node{
		data: value,
	}

	if queue.tail == nil && queue.head == nil {
		queue.tail = &newNode
		queue.head = &newNode
		return
	}

	queue.tail.next = &newNode
	queue.tail = &newNode
}

func (queue *Queue) Remove() interface{} {
	data := queue.head.data

	queue.head = queue.head.next

	return data
}