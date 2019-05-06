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

// 打印数组方法
template<typename T>
void printArray(T arr[], int n)
{
  for (int i = 0; i < n; i++)
    cout << arr[i] << " ";
  cout << endl;
  return;
}

template<typename T>
void selectionSort(T arr[], int n)
{
	for (int i = 0; i < n; i++)
	{
		// 寻找[i, n)区间里的最小值
		int minIndex = i;
		for (int j = i + 1; j < n; j++)
		{
			if (arr[j] < arr[minIndex])
				minIndex = j;
		}
    // 交换位置值
		swap(arr[i], arr[minIndex]);
	}
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

int main()
{
	int n = 10000;
	int* arr = generateRandomArray(n, 2, n);
  test("选择排序", selectionSort, arr, n);
	// printArray(arr, n);

	// 在helper中new过一个内存，需要将他释放掉
	delete[] arr;

	return 0;
}
