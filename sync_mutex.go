package main

import (
	"fmt"
	"sync"
	"time"
)

type signal struct{}

var ready bool
var mue sync.Mutex
var wg sync.WaitGroup

func workerMan(i int) {
	fmt.Printf("worker %d: is working...\n", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("worker %d: works done\n", i)
}

func spawnGroup(f func(i int), num int) <-chan signal {
	c := make(chan signal)

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for {
				mue.Lock()
				if !ready {
					mue.Unlock()
					time.Sleep(100 * time.Millisecond)
					continue
				}
				mue.Unlock()
				fmt.Printf("worker %d: start to work...\n", i)
				f(i)
				return
			}
		}(i + 1)
	}

	go func() {
		wg.Wait()
		c <- signal{}
		close(c)
	}()
	return c
}

func main() {
	fmt.Println("start a group of workers...")
	c := spawnGroup(workerMan, 5)

	time.Sleep(5 * time.Second) // 模拟 ready 前的准备工作
	fmt.Println("the group of workers start to work...")

	mue.Lock()
	ready = true
	mue.Unlock()

	<-c
	fmt.Println("the group of workers work done!")
}
