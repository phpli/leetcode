package main

import (
	"errors"
	"fmt"
	"github.com/lucasepe/codename"
	"github.com/phpli/tcp-server-demo1/frame"
	"github.com/phpli/tcp-server-demo1/packet"
	"net"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var num int = 2
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func(i int) {
			defer wg.Done()
			startClient(i)
		}(i + 1)
	}
	wg.Wait()
}

func startClient(i int) {
	quit := make(chan struct{})
	done := make(chan struct{})
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("dial error:", err)
		return
	}
	defer conn.Close()
	fmt.Printf("[client %d]: dial ok", i)
	rng, err := codename.DefaultRNG()
	if err != nil {
		//panic(err)
		fmt.Printf("[client %d]: get default rng error: %v", i, err)
	}
	frameCodec := frame.NewMyFrameCodec()
	var counter int
	go func() {
		for {
			select {
			case <-quit:
				done <- struct{}{}
				return
			default:
			}

			err = conn.SetReadDeadline(time.Now().Add(time.Second * 5))
			//if err != nil {
			//	return
			//}
			ackFramePayLoad, err := frameCodec.Decode(conn)
			fmt.Println(2222)
			if err != nil {
				var e net.Error
				if errors.As(err, &e) && e.Timeout() {
					continue
				}
				panic(err)
			}
			p, err := packet.Decode(ackFramePayLoad)
			submitAck, ok := p.(*packet.SubmitAck)
			if !ok {
				panic("not submitack")
			}
			fmt.Printf("[client %d]: the result of submit ack[%s] is %d\n", i, submitAck.ID, submitAck.Result)
		}
	}()

	for {
		counter++
		id := fmt.Sprintf("%d", counter)
		payload := codename.Generate(rng, 4)
		s := &packet.Submit{
			ID:      id,
			Payload: []byte(payload),
		}
		framePayload, err := packet.Encode(s)
		fmt.Printf(string(framePayload))
		if err != nil {
			panic(err)
		}
		fmt.Printf("[client %d]: send submit id = %s, payload=%s, frame length = %d\n",
			i, s.ID, s.Payload, len(framePayload)+4)
		err = frameCodec.Encode(conn, framePayload)
		if err != nil {
			panic(err)
		}
		time.Sleep(1 * time.Second)
		if counter >= 10 {
			quit <- struct{}{}
			<-done
			fmt.Printf("[client %d]: exit ok\n", i)
			return
		}
	}
}
