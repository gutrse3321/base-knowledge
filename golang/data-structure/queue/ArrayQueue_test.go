package queue

import (
	"fmt"
	"testing"
)

/**
 * @Author: Tomonori
 * @Date: 2020/2/18 15:02
 * @Title:
 * --- --- ---
 * @Desc:
 */
func TestArrayQueue(test *testing.T) {
	queue := NewArrayQueue()

	for i := 0; i < 10; i++ {
		queue.Enqueue(i)
		fmt.Println(queue)
	}
}

func TestDequeue(test *testing.T) {
	queue := NewArrayQueue()

	for i := 0; i < 10; i++ {
		queue.Enqueue(i)
		fmt.Println(queue)

		if i%3 == 2 {
			queue.Dequeue()
			fmt.Println(queue)
		}
	}
}
