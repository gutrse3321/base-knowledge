#pragma once
#include <iostream>
#include <ctime>
#include <cassert>

using namespace std;

namespace SortTestHelper
{
	// 生成随机数组
	int* generateRandomArray(int n, int rangeL, int rangeR)
	{
		// rangeL始终小于rangeR
		assert(rangeL <= rangeR);

		int* arr = new int[n];
		srand(time(NULL));

		// rand() 随机整数
		for (int i = 0; i < n; i++)
			arr[i] = rand() % (rangeR - rangeL + 1) + rangeL;
		return arr;
	}
}