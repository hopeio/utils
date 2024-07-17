package listqueue

type Node[T any] struct {
	data T
	prev *Node[T]
	next *Node[T]
}

func (n *Node[T]) Value() T {
	return n.data
}

func (n *Node[T]) Set(value T) {
	n.data = value
}

func (n *Node[T]) Previous() *Node[T] {
	return n.prev
}

func (n *Node[T]) Next() *Node[T] {
	return n.next
}

type ListQueue[T any] struct {
	head *Node[T]
	end  *Node[T]
	size int
	zero T
}

func New[T any](size int) *ListQueue[T] {
	q := &ListQueue[T]{}
	q.size = size
	return q
}

func (q *ListQueue[T]) Push(data T) {
	n := &Node[T]{data: data, next: nil}

	if q.end == nil {
		q.head = n
		q.end = n
	} else {
		q.end.next = n
		q.end = n
	}

	return
}

func (q *ListQueue[T]) Pop() (T, bool) {
	if q.head == nil {
		return q.zero, false
	}

	data := q.head.data
	q.head = q.head.next
	if q.head == nil {
		q.end = nil
	}

	return data, true
}
