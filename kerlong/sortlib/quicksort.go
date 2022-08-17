package sortlib

func QuickSort(array []int, start int, end int) {
	if start >= end {
		return
	}
	mid := array[start]
	i, j := start, end
	for i < j {
		//找到比mid小的
		for i < j && array[j] >= mid {
			j--
		}
		//找到比mid大的
		for i < j && array[i] <= mid {
			i++
		}
		if i < j {
			array[i], array[j] = array[j], array[i]
		}
	}
	//移动中间值
	array[j], array[start] = array[start], array[j]
	QuickSort(array, start, i)
	QuickSort(array, j+1, end)
}
