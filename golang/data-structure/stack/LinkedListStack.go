package stack

import (
	linkedList2 "all/linkedList"
	"bytes"
)

/**
 * @Author: Tomonori
 * @Date: 2020/2/20 14:56
 * @Title:
 * --- --- ---
 * @Desc:
 */

type LinkedListStack struct {
	linkedList *linkedList2.LinkedList
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{
		linkedList: linkedList2.NewLinkedList(),
	}
}

func (l *LinkedListStack) GetSize() int {
	return l.linkedList.GetSize()
}

func (l *LinkedListStack) IsEmpty() bool {
	return l.linkedList.IsEmpty()
}

func (l *LinkedListStack) Push(t interface{}) error {
	return l.linkedList.AddFirst(t)
}

func (l *LinkedListStack) Pop() (interface{}, error) {
	return l.linkedList.RemoveFirst()
}

func (l *LinkedListStack) Peek() (interface{}, error) {
	return l.linkedList.GetFirst()
}

func (l *LinkedListStack) String() string {
	buffer := bytes.Buffer{}
	buffer.WriteString("stack: top ")
	buffer.WriteString(l.linkedList.String())

	return buffer.String()
}
