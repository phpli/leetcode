package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

// trace.go

//func Trace(name string) func() {
//	println("enter:", name)
//	return func() {
//		println("exit:", name)
//	}
//}

func TraceTest() func() {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("not found caller")
	}
	fn := runtime.FuncForPC(pc)
	name := fn.Name()
	gid := curGoroutineID()
	fmt.Printf("g[%05d]: enter: [%s]\n", gid, name)
	return func() { fmt.Printf("g[%05d]: exit: [%s]\n", gid, name) }
}

//func foo() {
//	defer Trace()()
//	bar()
//}
//
//func bar() {
//	defer Trace()()
//}
//
//func main() {
//	defer Trace()()
//	foo()
//}

// trace2/trace.go
var goroutineSpace = []byte("goroutine ")

func curGoroutineID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	// Parse the 4707 out of "goroutine 4707 ["
	b = bytes.TrimPrefix(b, goroutineSpace)
	i := bytes.IndexByte(b, ' ')
	if i < 0 {
		panic(fmt.Sprintf("No space found in %q", b))
	}
	b = b[:i]
	n, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse goroutine ID out of %q: %v", b, err))
	}
	return n
}

// trace2/trace.go
func A1() {
	defer Trace()()
	B1()
}

func B1() {
	defer Trace()()
	C1()
}

func C1() {
	defer Trace()()
	D()
}

func D() {
	defer Trace()()
}

func A2() {
	defer Trace()()
	B2()
}
func B2() {
	defer Trace()()
	C2()
}
func C2() {
	defer Trace()()
	D()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		A2()
		wg.Done()
	}()

	A1()
	wg.Wait()
}

// trace3/trace.go

func printTrace(id uint64, name, arrow string, indent int) {
	indents := ""
	for i := 0; i < indent; i++ {
		indents += "    "
	}
	fmt.Printf("g[%05d]:%s%s%s\n", id, indents, arrow, name)
}

var mu sync.Mutex
var m = make(map[uint64]int)

func Trace() func() {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("not found caller")
	}

	fn := runtime.FuncForPC(pc)
	name := fn.Name()
	gid := curGoroutineID()

	mu.Lock()
	indents := m[gid]    // 获取当前gid对应的缩进层次
	m[gid] = indents + 1 // 缩进层次+1后存入map
	mu.Unlock()
	printTrace(gid, name, "->", indents+1)
	return func() {
		mu.Lock()
		indents := m[gid]    // 获取当前gid对应的缩进层次
		m[gid] = indents - 1 // 缩进层次-1后存入map
		mu.Unlock()
		printTrace(gid, name, "<-", indents)
	}
}
