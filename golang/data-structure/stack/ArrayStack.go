package stack

import (
	array2 "all/array/array"
	"bytes"
	"fmt"
)

func NewArrayStack() *ArrayStack {
	object := array2.NewArray()
	return &ArrayStack{array: object}
}

func NewArrayStackWithCap(cap int) *ArrayStack {
	object := array2.NewArrayWithCap(cap)
	return &ArrayStack{array: object}
}

type ArrayStack struct {
	array *array2.Array
}

func (a *ArrayStack) GetCap() int {
	return a.array.GetCap()
}

func (a *ArrayStack) GetSize() int {
	return a.array.GetSize()
}

func (a *ArrayStack) IsEmpty() bool {
	return a.array.IsEmpty()
}

func (a *ArrayStack) Push(t interface{}) error {
	return a.array.AddLast(t)
}

func (a *ArrayStack) Pop() (interface{}, error) {
	return a.array.RemoveLast()
}

func (a *ArrayStack) Peek() (interface{}, error) {
	return a.array.GetLast()
}

func (a *ArrayStack) String() string {
	buffer := bytes.Buffer{}
	buffer.WriteString("stack: ")
	buffer.WriteByte('[')

	for i := 0; i < a.GetSize(); i++ {
		if i < a.GetSize() {
			item, _ := a.array.Get(i)
			buffer.WriteString(fmt.Sprintf("%v", item))
			if i != a.GetSize()-1 {
				buffer.WriteString(", ")
			}
		}
	}

	buffer.WriteString("] on top")

	return buffer.String()
}
