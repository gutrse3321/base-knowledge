package queue

import (
	array2 "all/array/array"
	"bytes"
	"fmt"
)

/**
 * @Author: Tomonori
 * @Date: 2020/2/18 14:53
 * @Title:
 * --- --- ---
 * @Desc:
 */
type ArrayQueue struct {
	array array2.Array
}

func NewArrayQueue() *ArrayQueue {
	object := array2.NewArray()
	return &ArrayQueue{array: *object}
}

func NewArrayQueueWithCap(cap int) *ArrayQueue {
	object := array2.NewArrayWithCap(cap)
	return &ArrayQueue{array: *object}
}

func (a *ArrayQueue) GetCap() int {
	return a.array.GetCap()
}

func (a *ArrayQueue) GetSize() int {
	return a.array.GetSize()
}

func (a *ArrayQueue) IsEmpty() bool {
	return a.array.IsEmpty()
}

func (a *ArrayQueue) Enqueue(t interface{}) error {
	return a.array.AddLast(t)
}

func (a *ArrayQueue) Dequeue() (interface{}, error) {
	return a.array.RemoveFirst()
}

func (a *ArrayQueue) GetFront() (interface{}, error) {
	return a.array.GetFirst()
}

func (a *ArrayQueue) String() string {
	buffer := bytes.Buffer{}
	buffer.WriteString("Queue: ")
	buffer.WriteString("front [")

	for i := 0; i < a.GetSize(); i++ {
		if i < a.GetSize() {
			item, _ := a.array.Get(i)
			buffer.WriteString(fmt.Sprintf("%v", item))
			if i != a.GetSize()-1 {
				buffer.WriteString(", ")
			}
		}
	}

	buffer.WriteString("] tail")

	return buffer.String()
}
