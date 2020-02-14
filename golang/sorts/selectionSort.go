package main

import (
  Tools "golang/util"
)

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
  n := 100000
  arr, _ = Tools.GenerateRandomArray(n, 0, n)
  // 0.0682272s
  Tools.Test(arr, n, selectionSort, "选择")
}