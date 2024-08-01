package main

func case1() int {
	println("eval case1 expr")
	return 1
}

func case2_1() int {
	println("eval case2_1 expr")
	return 0
}
func case2_2() int {
	println("eval case2_2 expr")
	return 2
}

func case3() int {
	println("eval case3 expr")
	return 3
}

func switchexpr() int {
	println("eval switch expr")
	return 2
}

// 执行顺序，自上而下，自左而右
// break 只跳最内循环，要是跳循环 需要加label
func main() {
	switch switchexpr() {
	case case1():
		println("exec case1")
	case case2_1(), case2_2():
		println("exec case2")
	case case3():
		println("exec case3")
	default:
		println("exec default")
	}
}
