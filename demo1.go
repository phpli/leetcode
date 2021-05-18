package main

import "fmt"

//Two Sum

func main() {
	var nums = []int{2, 7, 11, 15}
	var target = 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		//fmt.Println(m[i])
		another := target - nums[i]
		//fmt.Println(another)
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
		//fmt.Println(m)
	}
	return nil
}
