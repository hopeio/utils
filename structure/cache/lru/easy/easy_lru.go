package lru

type Node[K comparable, V any] struct {
	Key        K
	Val        V
	Prev, Next *Node[K, V]
}
type LRUCache[K comparable, V any] struct {
	capacity   int
	data       map[K]*Node[K, V]
	Head, Tail *Node[K, V]
}

func New[K comparable, V any](capacity int) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		capacity: capacity,
		data:     make(map[K]*Node[K, V]),
		Head:     nil,
		Tail:     nil,
	}
}

func (l *LRUCache[K, V]) Get(key K) (V, bool) {
	v, ok := l.data[key]
	if ok {
		l.insert(v)
		return v.Val, true
	}
	return *new(V), false
}

func (l *LRUCache[K, V]) insert(v *Node[K, V]) {
	v.Prev.Next = v.Next
	v.Next.Prev = v.Prev
	head := l.Head
	head.Prev = v
	l.Head = v
	v.Next = head
}

func (l *LRUCache[K, V]) Put(key K, val V) {
	v, ok := l.data[key]
	if ok {
		l.insert(v)
		v.Val = val
	}

	if len(l.data) == l.capacity {
		delete(l.data, l.Tail.Key)
		l.Tail = l.Tail.Prev
	}

	newNode := &Node[K, V]{Key: key, Val: val}
	newNode.Next = l.Head
	l.Head.Prev = newNode
	l.Head = newNode
	l.data[key] = newNode
}
