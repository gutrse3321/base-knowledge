package linkedList

import (
	"fmt"
	"testing"
)

/**
 * @Author: Tomonori
 * @Date: 2020/2/19 18:26
 * @Title:
 * --- --- ---
 * @Desc:
 */

func TestSimple(t *testing.T) {
	linkedList := NewLinkedList()

	for i := 0; i < 5; i++ {
		linkedList.AddFirst(i)
		fmt.Println(linkedList)
	}

	linkedList.Add(3, 100)
	fmt.Println(linkedList)

	linkedList.Set(2, 50)
	fmt.Println(linkedList)

	fmt.Println(linkedList.Contains(1))

	fmt.Println(linkedList.Get(5))

	linkedList.Remove(5)
	fmt.Println(linkedList)

	linkedList.RemoveLast()
	fmt.Println(linkedList)
}
