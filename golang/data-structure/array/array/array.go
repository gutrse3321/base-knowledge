package array

import (
	"bytes"
	"errors"
	"fmt"
)

/**
 * @Author: Tomonori
 * @Date: 2020/2/14 12:27
 * @Title:
 * --- --- ---
 * @Desc:
 */
func NewArray() (object *Array) {
	return NewArrayWithCap(10)
}

func NewArrayWithCap(cap int) (object *Array) {
	object = &Array{
		data: make([]interface{}, cap),
		cap:  cap,
	}
	return
}

type Array struct {
	data []interface{}
	size int
	cap  int
}

func (a *Array) GetSize() int {
	return a.size
}

func (a *Array) GetCap() int {
	return len(a.data)
}

func (a *Array) Get(index int) (ele interface{}, err error) {
	err = a.requireIndexFailed(index)
	ele = a.data[index]
	return
}

func (a *Array) Set(index int, val interface{}) (err error) {
	err = a.requireIndexFailed(index)
	if err == nil {
		a.data[index] = val
	}
	return
}

func (a *Array) IsEmpty() bool {
	return a.size == 0
}

//FIXME Contains
//查找数组中是否有此元素
func (a *Array) Contains(val interface{}) bool {
	for ele := range a.data {
		if ele == val {
			return true
		}
	}
	return false
}

//FIXME Find
//查找数组元素的索引，-1为没找到
func (a *Array) Find(val interface{}) int {
	for index, item := range a.data {
		if item == val {
			return index
		}
	}
	return -1
}

//FIXME Add
//在第index的位置插入一个新元素
func (a *Array) Add(index int, val interface{}) (err error) {
	err = a.requireIndexFailed(index)

	if a.size == len(a.data) {
		a.resize(2 * len(a.data))
	}

	for i := a.size - 1; i >= index; i-- {
		a.data[i+1] = a.data[i]
	}

	a.data[index] = val
	a.size++

	return
}

//FIXME AddFirst
//所有元素前添加新元素
func (a *Array) AddFirst(val interface{}) (err error) {
	err = a.Add(0, val)
	return
}

//FIXME AddLast
//所有元素后添加新元素
func (a *Array) AddLast(val interface{}) (err error) {
	err = a.Add(a.size, val)
	return
}

//FIXME Remove
//从数组删除index索引的元素，返回删除的元素值
func (a *Array) Remove(index int) (val interface{}, err error) {
	err = a.requireIndexFailed(index)

	val = a.data[index]
	for i := index + 1; i < a.size; i++ {
		a.data[i-1] = a.data[i]
	}
	a.size--
	a.data[a.size] = nil

	//防止resize频繁，在1/4时才减容
	//不等于0防止数组只有一个元素的情况
	if a.size == len(a.data)/4 && len(a.data)/2 != 0 {
		a.resize(len(a.data) / 2)
	}

	return
}

//删除数组第一个元素
func (a *Array) RemoveFirst() (interface{}, error) {
	return a.Remove(0)
}

//删除数组最后一个元素
func (a *Array) RemoveLast() (interface{}, error) {
	return a.Remove(a.size - 1)
}

//FIXME RemoveElement
//删除数组中的某个元素
func (a *Array) RemoveElement(ele interface{}) {
	if index := a.Find(ele); index != -1 {
		a.Remove(index)
	}
}

func (a *Array) String() string {
	buffer := bytes.Buffer{}
	desc := fmt.Sprintf("size = %d, cap = %d\n", a.size, a.cap)
	buffer.WriteString(desc)
	buffer.WriteByte('[')

	for index, item := range a.data {
		if index < a.size {
			buffer.WriteString(fmt.Sprintf("%v", item))
			if index != a.size-1 {
				buffer.WriteString(", ")
			}
		}
	}

	buffer.WriteByte(']')
	buffer.WriteByte('\n')

	return buffer.String()
}

//FIXME resize
//处理数组容量
func (a *Array) resize(cap int) {
	newData := make([]interface{}, cap)

	for i := 0; i < a.size; i++ {
		newData[i] = a.data[i]
	}
	a.cap = len(newData)
	a.data = newData
}

//FIXME requireIndexFailed
//判断索引index是否超出数组元素数或小于0个数组元素数
func (a *Array) requireIndexFailed(index int) (err error) {
	if index < 0 || index >= a.size {
		return errors.New("func params index < 0 or index <= size")
	}
	return nil
}
