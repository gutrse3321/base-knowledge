package stack

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
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

func TestLinkedListStack(t *testing.T) {
	stack := NewLinkedListStack()

	for i := 0; i < 5; i++ {
		stack.Push(i)
		fmt.Println(stack)
	}

	stack.Pop()
	fmt.Println(stack)
}

func TestTime(t *testing.T) {
	count := 10000000

	arrayStack := NewArrayStack()
	fmt.Println(timeFun(arrayStack, count))

	linkedListStack := NewLinkedListStack()
	fmt.Println(timeFun(linkedListStack, count))
}

func timeFun(stack IStack, count int) float64 {
	startTime := time.Now()

	for i := 0; i < count; i++ {
		stack.Push(rand.Intn(math.MaxInt32))
	}

	for i := 0; i < count; i++ {
		stack.Pop()
	}

	endTime := time.Now()
	seconds := endTime.Sub(startTime).Seconds()
	return seconds
}
