package main

import "fmt"

//Container With Most Water

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}

func maxArea(height []int) int {
	start, end, max := 0, len(height)-1, 0
	for start < end {
		high := 0
		width := end - start
		if height[start] < height[end] {
			high = height[start]
			start++
		} else {
			high = height[end]
			end--
		}
		temp := width * high
		if temp > max {
			max = temp
		}
	}
	return max
}
