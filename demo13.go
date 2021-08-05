package main

import "fmt"

// 桶排序 适合于外部排序，就是内存比较小，数据比较大，数据需要放到磁盘上的那种。

func main()  {
	array := []int{31,16,37,2,13,32,10,27,7,42,29,18,28,12,9,}
	BucketSort(array)
	fmt.Println("BucketSort:",array)
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

func BucketSort(a []int) []int {
	num := len(a)
	if num <= 1{
		return a
	}
	max := getMax(a)
	buckets := make([][]int, num) // 二维切片
	index := 0
	for i := 0; i < num; i++ {
		index = a[i]*(num-1)/max // 桶序号
		buckets[index] = append(buckets[index],a[i])  // 加入对应的桶中
	}
	tmpPos := 0
	for i := 0; i < num; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0{
			sortInBucket(buckets[i])
			copy(a[tmpPos:], buckets[i])
			tmpPos += bucketLen
		}
	}
	return a
}

func sortInBucket(bucket []int)  { //这里可以使用任意一种排序
	length := len(bucket)
	if length <= 1{
		return
	}
	for i := 1;i<length;i++{
		backup := bucket[i]
		j := i-1
		for j>=0 && backup < bucket[j] {
			bucket[j+1] = bucket[j]
			j--
		}
		bucket[j+1] = backup
	}
}


