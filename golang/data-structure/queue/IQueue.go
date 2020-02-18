package queue

/**
 * @Author: Tomonori
 * @Date: 2020/2/18 14:51
 * @Title:
 * --- --- ---
 * @Desc:
 */
type IQueue interface {
	GetSize() int

	IsEmpty() bool

	Enqueue(t interface{}) error

	Dequeue() (interface{}, error)

	GetFront() (interface{}, error)
}
