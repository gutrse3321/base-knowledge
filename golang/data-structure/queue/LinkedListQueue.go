package queue

import (
	"bytes"
	"errors"
	"fmt"
)

/**
 * @Author: Tomonori
 * @Date: 2020/2/20 18:47
 * @Title:
 * --- --- ---
 * @Desc:
 */
type LinkedListQueue struct {
	head *node
	tail *node
	size int
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{
		head: nil,
		tail: nil,
	}
}

func (l *LinkedListQueue) GetSize() int {
	return l.size
}

func (l *LinkedListQueue) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedListQueue) Enqueue(t interface{}) error {
	if l.tail == nil {
		l.tail = newNode()
		l.head = l.tail
	} else {
		l.tail.next = newNodeWithVal(t)
		l.tail = l.tail.next
	}
	l.size++

	return nil
}

func (l *LinkedListQueue) Dequeue() (interface{}, error) {
	var err error
	if l.IsEmpty() {
		err = errors.New("Cannot dequeue from an empty queue")
	}

	deNode := l.head
	l.head = l.head.next
	deNode.next = nil

	if l.head == nil {
		l.tail = nil
	}
	l.size--

	return deNode, err
}

func (l *LinkedListQueue) GetFront() (interface{}, error) {
	var err error
	if l.IsEmpty() {
		err = errors.New("Cannot dequeue from an empty queue")
	}
	return l.head.val, err
}

func (l *LinkedListQueue) String() string {
	buffer := bytes.Buffer{}
	buffer.WriteString("Queue: ")
	buffer.WriteString("[head] ")

	current := l.head
	for current != nil {
		buffer.WriteString(fmt.Sprintf("%v -> ", current))
		current = current.next
	}
	buffer.WriteString(" nil [tail]")

	return buffer.String()
}

type node struct {
	val  interface{}
	next *node
}

func newNode() *node {
	return newNodeWithValAndNext(nil, nil)
}

func newNodeWithVal(val interface{}) *node {
	return newNodeWithValAndNext(val, nil)
}

func newNodeWithValAndNext(val interface{}, nde *node) *node {
	return &node{
		val:  val,
		next: nde,
	}
}

func (n *node) String() string {
	return fmt.Sprintf("%v", n.val)
}
