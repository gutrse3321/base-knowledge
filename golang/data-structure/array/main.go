package main

import (
	"array/array"
	"fmt"
)

/**
 * @Author: Tomonori
 * @Date: 2020/2/14 14:54
 * @Title:
 * --- --- ---
 * @Desc:
 */

func main() {
	arr := array.NewArray()

	for i := 0; i < 10; i++ {
		arr.AddLast(i)
	}
	fmt.Println(arr)

	fmt.Println(arr.Contains(-1))

	arr.AddFirst(-1)
	fmt.Println(arr)

	arr.Set(1, 100)
	arr.RemoveLast()
	fmt.Println(arr)

	arr.RemoveFirst()
	arr.RemoveFirst()
	arr.RemoveFirst()
	arr.RemoveFirst()
	arr.RemoveFirst()
	arr.RemoveFirst()
	arr.RemoveFirst()
	arr.RemoveFirst()
	fmt.Println(arr)
}
