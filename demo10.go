package main

import "fmt"

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