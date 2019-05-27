package main

import (
  Tools "golang/util"
  // "fmt"
)

func mergeSort(arr []int, n int) {
 mergeItem(arr, 0, n - 1)
}

func mergeItem(arr []int, l int, r int)  {
  if l >= r {
    return
  }
  mid := (l + r) / 2
  mergeItem(arr, l, mid)
  // fmt.Println(arr)
  mergeItem(arr, mid + 1, r)
  // fmt.Println(arr)
  if arr[mid] > arr[mid + 1] {
    merge(arr, l, mid, r)
  }
}

func merge(arr []int, l int, mid int, r int) {
  aux := make([]int, r - l + 1)
  for i := l; i <= r; i++ {
    aux[i - l] = arr[i]
  }

  i, j := l, mid + 1
  for k := l; k <= r; k++ {
    if i > mid {
      arr[k] = aux[j - l]
      j++
    } else if j > r {
      arr[k] = aux[i - l]
      i++
    } else if aux[i - l] < aux[j - l] {
      arr[k] = aux[i - l]
      i++
    } else {
      arr[k] = aux[j - l]
      j++
    }
  }
}

func main() {
  n := 100000
  arr, _ := Tools.GenerateRandomArray(n, 0, n)
  // mergeSort(arr, n)
  Tools.Test(arr, n, mergeSort, "归并")
}