package main

import (
	"fmt"
)

// 冒泡排序，比较相邻的俩个元素，看大小关系然后进行比较，通过一个中间变量换位置
// 插入排序，

func main() {
	values := []int{4, 93, 84, 85, 80, 37, 81, 94, 27, 12}
	//values[0]
	//fmt.Printf("%d",values[0])
	//fmt.Println(bubbleSort(values,len(values)))
	//fmt.Println(insertionSort(values,len(values)))
	//fmt.Println(selectSort(values,len(values)))
	secondSelectSort(values)
	//fmt.Printf("%d",len(values))
}

func bubbleSort(a []int, n int) []int {
	if n <= 1 {
		return a
	}
	for i := 0; i < n; i++ {
		var flag = false
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				var tmp = a[j]
				a[j] = a[j+1]
				a[j+1] = tmp
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return a
}

func copyAgain(a []int, n int) []int {
	if n <= 1 {
		return a
	}
	for i := 0; i < n; i++ {
		var flag = false
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				var tmp = a[j]
				a[j] = a[j+1]
				a[j+1] = tmp
				flag = true
			}
		}
		fmt.Printf("%d", a)
		fmt.Println("----")
		if !flag {
			break
		}
	}
	return a
}

/**
插入排序
*/
func insertionSort(a []int, n int) []int {
	if n <= 1 {
		return a
	}
	for i := 1; i < n; i++ {
		tmp := a[i]
		for j := i - 1; j >= 0; j-- {
			if tmp < a[j] {
				a[j+1] = a[j]
				a[j] = tmp
			} else {
				break
			}
		}
	}
	return a
}

func copyInsertion(a []int, n int) []int {
	if n <= 1 {
		return a
	}
	for i := 1; i < n; i++ {
		tmp := a[i]
		for j := i - 1; j >= 0; j-- {
			if a[j] > tmp {
				a[j+1] = a[j]
				a[j] = tmp
			} else {
				break
			}
		}
	}
	return a
}

func selectSort(a []int, n int) []int {
	if n <= 1 {
		return a
	}
	var index int
	for i := 0; i < n; i++ {
		index = i
		for j := i + 1; j < n; j++ {
			if a[index] > a[j] {
				index = j
			}
		}
		fmt.Printf("%d", a)
		fmt.Println("----")
		a[i], a[index] = a[index], a[i]
	}
	return a
}

func secondSelectSort(data []int) {
	length := len(data)
	for i := 0; i < length; i++ {
		tmp := data[i]
		flag := i
		for j := i + 1; j < length; j++ {
			if data[j] < tmp {
				tmp = data[j]
				flag = j
			}
		}
		if flag != i {
			data[flag] = data[i]
			data[i] = tmp
		}
		fmt.Println(data) //为了看具体排序的过程
	}
}
