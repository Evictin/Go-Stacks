package stack 

type denode struct {
	value interface{}
	next *denode
	prev *denode
}

type Queue struct {
	count int
	head *denode
	tail *denode
}

func (q Queue) Len() int {
	return q.count
}

func (q *Queue) Push(o interface{}) {
	node := &denode{o, q.head, q.tail}
	q.count++
	if q.head==nil {
		q.head, q.tail, node.prev, node.next = node, node, node, node
	} else {
		q.tail.next, q.head.prev, q.head = node, node, node
	}
}

func (q Queue) Peek() (value interface{}, ok bool) {
	if q.tail != nil {
		value, ok = q.tail.value, true
	}
	return 
}

func (q *Queue) Pop() (value interface{}, ok bool) {
	if q.tail != nil {
		value, ok = q.tail.value, true
		q.count--
		if q.count == 0 {
			q.head, q.tail = nil, nil
		} else {
			q.tail, q.head.prev, q.tail.prev.next = q.tail.prev, q.tail.prev, q.head
		}
	}
	return
}

func (q Queue) IsEmpty() bool {
	return q.count == 0
}
