package main

type Stack interface {
	GetSize() int
	IsEmpty() bool
	Push(t interface{}) error
	Pop() (interface{}, error)
	Peek() (interface{}, error)
}
