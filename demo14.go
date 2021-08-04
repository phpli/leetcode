package main

import "fmt"

/**
计数排序是桶排序的特殊变种
 */

func main()  {
	array := []int{69,69,16,48,2,2,3,3,32,10,27,17,42,29,8,28,12,9,}
	countingSort(array)
	fmt.Println("BucketSort:",array)
}

func countingSort(array []int) []int{
	arrayLen := len(array)
	if arrayLen <= 1 {
		return array
	}
	maxValue := getMax(array)
	bucketLen := maxValue + 1
	bucket := make([]int,bucketLen) // 初始为0的数组
	sortIndex := 0

	for i := 0; i<arrayLen; i++ {
		bucket[array[i]] += 1
	}
	fmt.Println(bucket)
	for j := 0; j < bucketLen; j++ {
		for bucket[j] > 0 {
			array[sortIndex] = j
			sortIndex += 1
			bucket[j] -= 1
		}
	}
	return array
}

func getMax(a []int) int  {
	max := a[0]
	for i := 0; i < len(a) ; i++  {
		if a[i] > max{
			max = a[i]
		}
	}
	return max
}
