#include "pch.h"
#include <iostream>
#include "SortTestHelper.h"

using namespace std;

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

template<typename T>
void printArray(T arr[], int n)
{
	for (int i = 0; i < n; i++)
		cout << arr[i] << " ";
	cout << endl;
	return;
}

int main()
{
	int n = 10;
	int* arr = SortTestHelper::generateRandomArray(n, 2, n);
	selectionSort(arr, n);

	// 打印
	printArray(arr, n);
	// 在helper中new过一个内存，需要将他释放掉
	delete[] arr;

	return 0;
}
