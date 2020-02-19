package queue

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

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

func TestTime(t *testing.T) {
	fmt.Println(math.MaxInt32)

	count := 100000

	arrayQueue := NewArrayQueue()
	fmt.Println(timeFun(arrayQueue, count).Seconds())

	loopQueue := NewLoopQueue()
	fmt.Println(timeFun(loopQueue, count).Seconds())
}

func timeFun(queue IQueue, count int) time.Duration {
	startTime := time.Now()

	for i := 0; i < count; i++ {
		queue.Enqueue(rand.Intn(math.MaxInt32))
	}

	for i := 0; i < count; i++ {
		queue.Dequeue()
	}

	endTime := time.Now()

	return endTime.Sub(startTime)
}
