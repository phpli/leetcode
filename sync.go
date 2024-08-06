package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // 在函数结束时调用 Done()，表示 goroutine 完成
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second * 2) // 模拟工作
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // 增加等待计数器的值
		go worker(i, &wg)
	}

	wg.Wait() // 阻塞，直到所有 goroutine 完成
	fmt.Println("All workers done")
}
