package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	//fmt.Println(threeSum(nums))
	fmt.Println(sumThree(nums))
	//fmt.Println(threeSum2(nums))
}

//用 map 提前计算好任意 3 个数字之和，保存起来，可以将时间复杂度降到 O(n^2)。这一题比较麻烦的一点在于，
//最后输出解的时候，要求输出不重复的解。数组中同一个数字可能出现多次，同一个数字也可能使用多次，
//但是最后输出解的时候，不能重复。例如 [-1，-1，2] 和 [2, -1, -1]、[-1, 2, -1]
//这 3 个解是重复的，即使 -1 可能出现 100 次，每次使用的 -1 的数组下标都是不同的。
//这里就需要去重和排序了。map 记录每个数字出现的次数，然后对 map 的 key 数组进行排序，最后在这个排序以后的数组里面扫，找到另外 2 个数字能和自己组成 0 的组合

func threeSum(nums []int) [][]int {

	sort.Ints(nums)
	result, start, end, index, addNum, length := make([][]int, 0), 0, 0, 0, 0, len(nums)

	for index = 1; index < length-1; index++ {
		start, end = 0, length-1
		if index > 1 && nums[index] == nums[index-1] {
			start = index - 1
		}
		for start < index && end > index {
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < length-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			addNum = nums[start] + nums[end] + nums[index]
			if addNum == 0 {
				result = append(result, []int{nums[start], nums[index], nums[end]})
				start++
				end--
			} else if addNum > 0 {
				end--
			} else {
				start++
			}
		}
	}
	return result
}

func threeSum1(nums []int) [][]int {
	res := [][]int{}
	counter := map[int]int{}
	for _, value := range nums {
		counter[value]++
	}
	fmt.Println(counter)
	uniqNums := []int{}
	for key := range counter {
		uniqNums = append(uniqNums, key)
	}
	fmt.Println(uniqNums)
	sort.Ints(uniqNums)
	for i := 0; i < len(uniqNums); i++ {
		if (uniqNums[i]*3 == 0) && counter[uniqNums[i]] >= 3 {
			res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[i]})
		}
		for j := i + 1; j < len(uniqNums); j++ {
			if (uniqNums[i]*2+uniqNums[j] == 0) && counter[uniqNums[i]] > 1 {
				res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j]})
			}
			if (uniqNums[j]*2+uniqNums[i] == 0) && counter[uniqNums[j]] > 1 {
				res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[j]})
			}
			c := 0 - uniqNums[i] - uniqNums[j]
			if c > uniqNums[j] && counter[c] > 0 {
				res = append(res, []int{uniqNums[i], uniqNums[j], c})
			}
		}
	}
	return res
}

func threeSum2(nums []int) [][]int {
	res := [][]int{}
	counter := map[int]int{}
	for _, value := range nums {
		counter[value]++
	}

	uniqNums := []int{}
	for key := range counter {
		uniqNums = append(uniqNums, key)
	}
	sort.Ints(uniqNums)

	for i := 0; i < len(uniqNums); i++ {
		if (uniqNums[i]*3 == 0) && counter[uniqNums[i]] >= 3 {
			res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[i]})
		}
		for j := i + 1; j < len(uniqNums); j++ {
			if (uniqNums[i]*2+uniqNums[j] == 0) && counter[uniqNums[i]] > 1 {
				res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j]})
			}
			if (uniqNums[j]*2+uniqNums[i] == 0) && counter[uniqNums[j]] > 1 {
				res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[j]})
			}
			c := 0 - uniqNums[i] - uniqNums[j]
			if c > uniqNums[j] && counter[c] > 0 {
				res = append(res, []int{uniqNums[i], uniqNums[j], c})
			}
		}
	}
	return res
}

func sumThree(nums []int) [][]int {

	sort.Ints(nums)
	result, start, end, index, addNum, length := make([][]int, 0), 0, 0, 0, 0, len(nums)
	for index = 1; index < length; index++ {
		start, end = 0, length-1
		//对起始位置赋值
		if index > 1 && nums[index] == nums[index-1] {
			start = index - 1
		}
		//循环的起始条件
		for start < index && end > index {
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < length-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			addNum = nums[start] + nums[end] + nums[index]
			if addNum == 0 {
				result = append(result, []int{nums[start], nums[index], nums[end]})
				start++
				end++
			} else if addNum > 0 {
				end--
			} else {
				start++
			}
		}
	}
	return result
}
