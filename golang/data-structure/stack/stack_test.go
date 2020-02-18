package main

import (
	"fmt"
	"testing"
)

func TestStack(test *testing.T) {
	fmt.Println("isValid:", isValid("([])"))

	stack := NewArrayStack()

	for i := 0; i < 5; i++ {
		stack.Push(i)
		fmt.Println(stack)
	}

	stack.Pop()
	fmt.Println(stack)
}
