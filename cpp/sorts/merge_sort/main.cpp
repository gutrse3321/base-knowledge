#include <iostream>
#include <cstdlib>
#include <ctime>
#include <cassert>

using namespace std;

// 生成随机元素数组
int* generateRandomArray(int n, int rangeL, int rangeR)
{
	// rangeL始终小于或等于rangeR
	assert(rangeL <= rangeR);

	int* arr = new int[n];
  // 随机种子设置，当前时间
	srand(time(NULL));

	// rand() 随机整数
	for (int i = 0; i < n; i++)
		arr[i] = rand() % (rangeR - rangeL + 1) + rangeL;
	return arr;
}

// 判断是否是正确的排序结果
template<typename T>
bool isSorted(T arr[], int n)
{
  for (int i = 0; i < n - 1; i++)
    if (arr[i] > arr[i + 1])
      return false;
  return true;
}

// 计算排序运行时间
template<typename T>
void test(string sortName, void(*sort)(T[], int), T arr[], int n)
{
  clock_t startTime = clock();
  sort(arr, n);
  clock_t endTime = clock();

  assert(isSorted(arr, n));

  cout << sortName << ": " << double(endTime - startTime) / CLOCKS_PER_SEC << "s" << endl;
  return;
}

// 将arr[l...mid]和arr[mid + 1...r]两部分进行归并
template<typename T>
void merge(T arr[], int l, int mid, int r)
{
  T aux[r - l + 1];
  for (int i = l; i <= r; i++)
    aux[i - l] = arr[i];
  
  int i = l, j = mid + 1;
  for (int k = l; k <= r; k++)
  {
    if (i > mid)
    {
      arr[k] = aux[j - l];
      j++;
    }
    else if (j > r)
    {
      arr[k] = aux[i - l];
      i++;
    }
    else if (aux[i - l] < aux[j - l])
    {
      arr[k] = aux[i - l];
      i++;
    }
    else
    {
      arr[k] = aux[j - l];
      j++;
    }
  }
}

// 递归使用归并排序，对arr[l...r]的范围进行排序
template<typename T>
void mergeItem(T arr[], int l, int r)
{
  if (l >= r)
    return;
  int mid = (l + r) / 2;
  mergeItem(arr, l, mid);
  mergeItem(arr, mid + 1, r);
  merge(arr, l, mid, r);
}

template<typename T>
void MergeSort(T arr[], int n)
{
  mergeItem(arr, 0, n - 1);
}

int main()
{
  int n = 10000;
	int* arr = generateRandomArray(n, 2, n);
  // 0.001934s
  test("cpp 归并排序", MergeSort, arr, n);

  delete[] arr;

  return 0;
}
