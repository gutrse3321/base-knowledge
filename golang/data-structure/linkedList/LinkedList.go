package linkedList

import (
	"bytes"
	"errors"
	"fmt"
)

/**
 * @Author: Tomonori
 * @Date: 2020/2/19 17:32
 * @Title:
 * --- --- ---
 * @Desc:
 */

type LinkedList struct {
	dummyHead *node
	size      int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		dummyHead: newNodeWithValAndNext(nil, nil),
	}
}

func (l *LinkedList) GetSize() int {
	return l.size
}

func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

/**
 * TODO 根据索引位置添加节点(不常用)
 * 在链表的index位置添加新的元素
 * @param index
 * @param t
 */
func (l *LinkedList) Add(index int, t interface{}) (err error) {
	if index < 0 || index > l.size {
		err = errors.New("Add failed, Illegal index")
	}

	prev := l.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}

	//找到index - 1的节点元素，构造一个新的节点，并将index之前的
	//下一个节点赋给新的节点的next，最后将新的节点赋值index的next
	prev.next = newNodeWithValAndNext(t, prev.next)
	l.size++
	return
}

func (l *LinkedList) AddFirst(t interface{}) error {
	return l.Add(0, t)
}

func (l *LinkedList) AddLast(t interface{}) error {
	return l.Add(l.size, t)
}

func (l *LinkedList) Get(index int) (t interface{}, err error) {
	if index < 0 || index >= l.size {
		err = errors.New("Get failed, Illegal index")
	}

	current := l.dummyHead.next
	for i := 0; i < index; i++ {
		current = current.next
	}
	t = current.val

	return
}

func (l *LinkedList) GetFirst() (interface{}, error) {
	return l.Get(0)
}

func (l *LinkedList) GetLast() (interface{}, error) {
	return l.Get(l.size - 1)
}

func (l *LinkedList) Set(index int, t interface{}) (err error) {
	if index < 0 || index >= l.size {
		err = errors.New("Set failed, Illegal index")
	}

	current := l.dummyHead.next
	for i := 0; i < index; i++ {
		current = current.next
	}
	current.val = t

	return
}

func (l *LinkedList) Contains(t interface{}) bool {
	current := l.dummyHead.next

	for current != nil {
		if current.val == t {
			return true
		}
		current = current.next
	}

	return false
}

//移除索引的节点
func (l *LinkedList) Remove(index int) (t interface{}, err error) {
	if index < 0 || index >= l.size {
		err = errors.New("Remove failed, Index is Illegal")
	}

	prev := l.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}

	removeNode := prev.next
	prev.next = removeNode.next
	removeNode.next = nil
	l.size--

	return removeNode.val, err
}

func (l *LinkedList) RemoveFirst() (interface{}, error) {
	return l.Remove(0)
}

func (l *LinkedList) RemoveLast() (interface{}, error) {
	return l.Remove(l.size - 1)
}

func (l *LinkedList) String() string {
	buffer := bytes.Buffer{}

	for current := l.dummyHead.next; current != nil; current = current.next {
		buffer.WriteString(fmt.Sprintf("%v -> ", current.val))
	}

	buffer.WriteString("nil")

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
