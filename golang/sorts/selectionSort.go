package main

import (
  "errors"
  "time"
  "math/rand"
  "fmt"
)

func init() {
  rand.Seed(time.Now().UnixNano())
}

func GenerateRandomArray(n, rangeL, rangeR int) ([]int, error) {
  var err error
  if rangeL > rangeR {
    err = errors.New("rangeL始终小于或等于rangeR")
    fmt.Println(err)
    return nil, err
  }

  arr := make([]int, n)
  for i := 0; i < n; i++ {
    temp := rand.Int() % (rangeR - rangeL + 1) + rangeL
    arr[i] = temp
  }
  return arr, nil
}

func selectionSort(arr []int, n int) {
  for i := 0; i < n; i++ {
		// 寻找[i, n)区间里的最小值
		minIndex := i;
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j;
      }
		}
    // 交换位置值
		arr[i], arr[minIndex] = arr[minIndex], arr[i];
	}
}

func main() {
  var arr []int
  n := 10
  arr, _ = GenerateRandomArray(n, 0, n)
  fmt.Println(arr)
  selectionSort(arr, n)
  fmt.Println(arr)
}