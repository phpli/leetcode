package main

import (
	"fmt"
	"sort"
)

func main()  {
	array := []int{3,69,69,16,48,2,3,3,32,32,10,27,27,17,42,29,8,28,12,9}
	sort.Ints(array)
	fmt.Println("BucketSort:",array)
	index := firstValueBinarySearch(array,len(array),27)
	index2 := lastValueBinarySearch(array,len(array),27)
	index3 := firstGtValueBinarySearch(array,len(array),27)
	index4 := firstLtValueBinarySearch(array,len(array),27)
	fmt.Println("BinarySearch:",index)
	fmt.Println("BinarySearch:",index2)
	fmt.Println("BinarySearch:",index3)
	fmt.Println("BinarySearch:",index4)
}

/**
变体一：查找第一个值等于给定值的元素
 */
func firstValueBinarySearch(array []int, n,value int) int {
	low := 0
	high := n - 1
	for low <= high {
		mid := low + ((high-low)>>1)
		if array[mid] > value  {
			high = mid - 1
		}else if array[mid] < value {
			low = mid + 1
		}else {
			if mid == 0 || array[mid-1] != value{
					return mid
			}else{
				high = mid - 1
			}
		}
	}
	return -1
}


/**
变体二：查找最后一个值等于给定值的元素
 */
func lastValueBinarySearch(array []int, n,value int) int {
	low := 0
	high := n - 1
	for low <= high {
		mid := low + ((high-low)>>1)
		if array[mid] > value{
			high = mid - 1
		}else if array[mid] < value {
			low = mid + 1
		}else {
			if mid == n-1 || array[mid+1] != value  { // 这已经是队尾
				return mid
			}else {
				low = mid + 1
			}
		}
	}
	return -1
}



/**
变体三：查找第一个大于等于给定值的元素
 */
func firstGtValueBinarySearch(array []int, n,value int) int {
	low := 0
	high := n - 1
	for low <= high {
		mid := low + ((high-low)>>1)
		if array[mid] >= value{
			if mid==0 || array[mid-1] < value{
				return mid
			}else {
				high = mid - 1
			}
		}else {
			low = mid + 1
		}
	}
	return -1
}

/**
变体四：查找最后一个小于等于给定值的元素
 */
func firstLtValueBinarySearch(array []int, n,value int) int {
	low := 0
	high := n-1
	for low <= high  {
		mid := low + ((high-low)>>1)
		if array[mid] <= value{
			if mid == n-1 || array[mid+1] > value {
				return mid
			}else {
				low = mid + 1
			}
		}else {
			high = mid - 1
		}
	}
	return -1
}

