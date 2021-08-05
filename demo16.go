package main

import (
	"fmt"
	"sort"
)

/**
二分法查找：适用于有序的数据，且适用于底层结构为 内存连续的数组数据，不连续为链表
首先，二分查找依赖的是顺序表结构，简单点说就是数组
其次，二分查找针对的是有序数据。
再次，数据量太小不适合二分查找。
最后，数据量太大也不适合二分查找。
 */

func main()  {
	array := []int{69,16,48,2,3,32,10,27,17,42,29,8,28,12,9}
	sort.Ints(array)
	fmt.Println("BucketSort:",array)
	index := bsearch(array,len(array),27)
	fmt.Println("BucketSort:",index)
}


func bsearch(array []int, n,value int) int{
	low := 0
	high := n - 1
	for low <= high {
		mid := low + ((high -low)>>1)
		if array[mid] == value {
			return mid
		}else if array[mid] < value {
			low = mid + 1
		}else if array[mid] > value{
			high = mid - 1
		}
	}
	return -1
}






