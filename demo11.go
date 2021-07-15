package main

import "fmt"

/**
快排排序
如果要排序数组中下标从 p 到 r 之间的一组数据，
我们选择 p 到 r 之间的任意一个数据作为 pivot（分区点）。
我们遍历 p 到 r 之间的数据，将小于 pivot 的放到左边，
将大于 pivot 的放到右边，
将 pivot 放到中间。
经过这一步骤之后，
数组 p 到 r 之间的数据就被分成了三个部分，前面 p 到 q-1 之间都是小于 pivot 的，
中间是 pivot，
后面的 q+1 到 r 之间是大于 pivot 的。
根据分治、递归的处理思想，我们可以用递归排序下标从 p 到 q-1 之间的数据和下标从 q+1 到 r 之间的数据，直到区间缩小为 1，就说明所有的数据都有序了。
*/
//排序

func main() {
	var arr = []int{-1, 5, 2, 1, 4, 3}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
func quickSort(data []int, start, end int) {
	var pivotPos int
	if start < end {
		pivotPos = partition(data, start, end)
		quickSort(data, start, pivotPos-1)
		quickSort(data, pivotPos+1, end)
	}
}


//划定分区点
func partition(data []int, start, end int) int {
	pivot := data[start]
	for start < end {
		for end > start && data[end] > pivot {
			end--
		}
		if end > start {
			data[start] = data[end]
			start++
		}
		for start < end && data[start] < pivot {
			start++
		}
		if start < end {
			data[end] = data[start]
			end--
		}
	}
	data[start] = pivot
	return start
}
