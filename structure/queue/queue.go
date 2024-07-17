package queue

type Queue[T any] interface {
	// 获取当前链表长度。
	Len() int
	// 获取当前链表容量。
	Capacity() int
	// 获取当前链表头结点。
	Front() (T, bool)
	// 获取当前链表尾结点。
	Tail() (T, bool)
	// 入列。
	Enqueue(value T) bool
	// 出列。
	Dequeue() T
}
