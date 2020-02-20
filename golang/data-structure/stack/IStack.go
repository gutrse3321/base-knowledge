package stack

type IStack interface {
	GetSize() int
	IsEmpty() bool
	Push(t interface{}) error
	Pop() (interface{}, error)
	Peek() (interface{}, error)
}
