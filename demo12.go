package main

import "fmt"

/**
所谓分区就是找到第一个元素的终极位置
 */
func main()  {
	var arr = []int{0, 5, 2, 1, 4, 3}
	//part(arr, 0, len(arr)-1)
	 k := findKthLargest(arr,3)
	fmt.Println(k)
}

func part(data []int, start,end int) int  {
	v := data[start]
	less := start
	for i := start+1; i < end; i++ {
		if  v < data[i]{
			break
		}
		less++
		if i != less{
			data[less],data[i] = data[i],data[less]
		}
	}
	data[less], data[start] = data[start], data[less]
	return less
}

func findKthLargest(data []int,k int) int  {
	start := 0
	n := len(data)
	index := n - k
	for  {
		p := part(data,start,n)
		//fmt.Println(p)
		if index == p{
			return data[p]
		}else if index < p { // value在[l...p)里面
			n = p
		} else { // value在[p+1...r)里面
			start = p + 1
		}
	}
}