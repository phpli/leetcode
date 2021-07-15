package main

import "fmt"

/**
分治思想
归并算法思想：
1.把数组分钟2段
2.拿一个临时数组
3.都在i,j 都小于各自的数组长度时拿第一段里的第一个数据i和第二段里的第一个数据j比较
小于的一个放入到临时数组，并且把小的一个索引+1
4.判断i,j 小于数组的大小的时候，依次放到临时数组里
 */
/**
归并排序是稳定算法，排序过程中可以保持部分数据位置不变动
时间复杂度：O(nlogn)
空间复杂度：O(n) 归并排序不是原地排序算法。
 */
func main()  {
	fmt.Println("归并排序学习与演示")
	array := []int{9,8,7,6,5,4,3,3,2,1}
	//show(array)
	newArray := mergeSort(array)
	show(newArray)
}

//归
func mergeSort(array []int) []int  {
	if len(array)<2{
		return array
	}
	mid := len(array)/2
	left := mergeSort(array[:mid])
	right := mergeSort(array[mid:])
	result := merge(left, right)
	return result
}

//并
func merge(left, right []int) []int {
	fmt.Println(left,right)
	temp := make([]int,0)
	i,j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j]{
			temp = append(temp,left[i])
			i++
		}else {
			temp = append(temp,right[j])
			j++
		}
	}
	if i < len(left){
		temp = append(temp,left[i:]...)
	}

	if j < len(right){
		temp = append(temp,right[j:]...)
	}
	return temp
}

func show(arr []int)  {
	for _,item := range arr{
		fmt.Printf("%d",item)
	}
	fmt.Println("")
}