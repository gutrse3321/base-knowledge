#pragma once
#include <iostream>
#include <ctime>
#include <cassert>

using namespace std;

namespace SortTestHelper
{
	// ����n��Ԫ�ص�������飬ÿ��Ԫ�ص������ΧΪrangeL��rangeR
	int* generateRandomArray(int n, int rangeL, int rangeR)
	{
		// rangeLҪС��rangeR
		assert(rangeL <= rangeR);

		int* arr = new int[n];
		srand(time(NULL));

		// rand() �����������
		for (int i = 0; i < n; i++)
			arr[i] = rand() % (rangeR - rangeL + 1) + rangeL;
		return arr;
	}
}