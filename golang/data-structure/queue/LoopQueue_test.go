package queue

import (
	"fmt"
	"testing"
)

/**
 * @Author: Tomonori
 * @Date: 2020/2/18 15:42
 * @Title:
 * --- --- ---
 * @Desc:
 */
func TestNormalLoopQueue(t *testing.T) {
	queue := NewLoopQueue()
	for i := 0; i < 10; i++ {
		queue.Enqueue(i)
		fmt.Println(queue)
	}
}

func TestLoopQueue_Dequeue(t *testing.T) {
	queue := NewLoopQueueWithCap(10)
	for i := 0; i < 10; i++ {
		queue.Enqueue(i)
		fmt.Println(queue)

		if i%3 == 2 {
			queue.Dequeue()
			fmt.Println(queue)
		}
	}
}
