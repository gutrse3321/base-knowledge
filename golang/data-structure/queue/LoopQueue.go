package queue

import (
	"bytes"
	"errors"
	"fmt"
)

/**
 * @Author: Tomonori
 * @Date: 2020/2/18 15:05
 * @Title: 循环队列
 * --- --- ---
 * @Desc: eg: [空, front, ..., tail], [空, 空, front, ... , tail, 空,空], [tail, 空, front,...]
 */

//循环队列，最后tail始终留一个空位
//tail为最后一个元素的下一位索引
type LoopQueue struct {
	data  []interface{}
	front int
	tail  int
	size  int
}

func NewLoopQueue() *LoopQueue {
	return NewLoopQueueWithCap(10)
}

//FIXME 循环队列，最后tail始终留一个空位
//所以 +1
func NewLoopQueueWithCap(cap int) *LoopQueue {
	return &LoopQueue{
		data: make([]interface{}, cap+1),
	}
}

func (l *LoopQueue) GetCap() int {
	return len(l.data) - 1
}

func (l *LoopQueue) GetSize() int {
	return l.size
}

func (l *LoopQueue) IsEmpty() bool {
	return l.front == l.tail
}

//FIXME Enqueue 添加队列元素
func (l *LoopQueue) Enqueue(t interface{}) error {
	//一个循环，如果tail在front之前，需要求余获取位置
	//判断是否和front相同，则扩容
	if (l.tail+1)%len(l.data) == l.front {
		l.resize(l.GetCap() * 2)
	}

	l.data[l.tail] = t
	l.tail = (l.tail + 1) % len(l.data)
	l.size++

	return nil
}

//FIXME 移除头部
func (l *LoopQueue) Dequeue() (interface{}, error) {
	err := l.panicEmptyQueueError()

	headEle := l.data[l.front]
	l.data[l.front] = nil
	//将头部索引移至上一个头部的下一位
	l.front = (l.front + 1) % len(l.data)
	l.size--

	//当队列元素数量和数组容量相同，且容量不等于两个
	if l.size == l.GetCap()/4 && l.GetCap()/2 != 0 {
		l.resize(l.GetCap() / 2)
	}

	return headEle, err
}

func (l *LoopQueue) GetFront() (interface{}, error) {
	err := l.panicEmptyQueueError()
	return l.data[l.front], err
}

func (l *LoopQueue) resize(cap int) {
	newData := make([]interface{}, cap+1)

	for i := 0; i < l.size; i++ {
		//根据front+i和数组总数求余得到当前front和其他元素的索引
		newData[i] = l.data[(l.front+i)%len(l.data)]
	}

	//新的数组从零开始，则front和tail重新赋值
	l.data = newData
	l.front = 0
	l.tail = l.size
}

func (l *LoopQueue) panicEmptyQueueError() (err error) {
	if l.IsEmpty() {
		err = errors.New("error  from an empty queue.")
	}
	return
}

func (l *LoopQueue) String() string {
	buffer := bytes.Buffer{}
	desc := fmt.Sprintf("size = %d, cap = %d\n", l.size, l.GetCap())
	buffer.WriteString(desc)
	buffer.WriteString("front [")

	//从头部开始遍历，等于尾部结束
	for i := l.front; i != l.tail; i = (i + 1) % len(l.data) {
		buffer.WriteString(fmt.Sprintf("%v", l.data[i]))
		if (i+1)%len(l.data) != l.tail {
			buffer.WriteString(", ")
		}
	}

	buffer.WriteString("] tail\n")

	return buffer.String()
}
