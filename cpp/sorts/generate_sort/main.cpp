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
		swap(arr[i], arr[minIndex]);
	}
}

int main()
{
	int n = 10;
	int* arr = generateRandomArray(n, 2, n);
	selectionSort(arr, n);
	printArray(arr, n);

	// 在helper中new过一个内存，需要将他释放掉
	delete[] arr;

	return 0;
}
