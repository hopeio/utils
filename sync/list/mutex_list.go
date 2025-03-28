package list

import (
	"sync"
	"sync/atomic"
)

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type MutexList[T any] struct {
	head *Node[T]
	tail *Node[T]
	mu   sync.RWMutex
	size uint64
	zero T
}

func NewMutexList[T any]() *MutexList[T] {
	return &MutexList[T]{}
}

func (l *MutexList[T]) Push(v T) {
	l.mu.Lock()
	defer l.mu.Unlock()
	node := &Node[T]{v, nil}
	if l.size == 0 {
		l.head = node
		l.tail = node
		l.size++
		return
	}
	l.tail.Next = node
	l.tail = node
	l.size++
}

func (l *MutexList[T]) Pop() (v T, ok bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.size == 0 {
		return l.zero, false
	}

	p := l.head
	l.head = p.Next
	if l.size == 1 {
		l.tail = nil
	}
	l.size--
	return p.Value, true
}

func (l *MutexList[T]) Len() uint64 {
	return atomic.LoadUint64(&l.size)
}
