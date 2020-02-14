package util

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

type sorts func([]int, int)

func Test(arr []int, n int, sort sorts, str string) {
  start := time.Now()
  sort(arr, n)
  end := time.Now()
  fmt.Println("golang ", str, "排序: ", end.Sub(start))
}
