package main

import (
	"fmt"
	"github.com/codingWhat/go-cowboy/frame"
	"github.com/codingWhat/go-cowboy/packet"
	"net"
	"sync"
)

func main() {
	//conn, err := net.Dial("tcp", ":8888")
	//if err != nil {
	//	panic(err)
	//	return
	//}
	//
	//codec := &frame.MyFrameCodec{}
	//pack := packet.Submit{ID: fmt.Sprintf("%08d", 1), Payload: []byte("absolute-karatecha")}
	//payload, err := pack.Encode()
	//if err != nil {
	//	panic(err)
	//}
	//err = codec.Encode(conn, payload)
	//if err != nil {
	//	panic(err)
	//	return
	//}
	//
	//ret, err := codec.Decode(conn)
	//if err != nil {
	//	panic(err)
	//	return
	//}
	//
	//fmt.Println("has received:", string(ret), err)

	wg := &sync.WaitGroup{}

	for i := 0; i < 500; i++ {
		wg.Add(1)
		go startClient(i+1, wg)
	}

	wg.Wait()

}

func startClient(i int, wg *sync.WaitGroup) {
	defer wg.Done()

	var (
		//quit        = make(chan struct{})
		//done        = make(chan struct{})
		counter int = 1
	)

	//conn, err := net.Dial("tcp", ":8888")
	conn, err := net.Dial("tcp", "9.134.199.174:8888")
	if err != nil {
		panic(err)
		return
	}
	defer func() {
		conn.Close()
		fmt.Println("client:", i, " 退出")
	}()

	codec := frame.NewMyFrameCodec()

	go func() {
		for {
			//select {
			//case <-quit:
			//	done <- struct{}{}
			//	return
			//default:
			//
			//}

			body, err := codec.Decode(conn)
			if err != nil {
				panic(err)
				return
			}

			_, err = packet.Decode(body)
			if err != nil {
				panic(err)
			}

			//fmt.Printf("client: %d---receive, %+v \n", i, cmd)
		}

	}()

	for {

		submit := &packet.Submit{
			ID:      fmt.Sprintf("%08d", counter),
			Payload: []byte("hello"),
		}

		payload, err := packet.Encode(submit)
		if err != nil {
			panic(err)
		}
	//	fmt.Println("client:", i, "---send, ", string(payload))
		err = codec.Encode(conn, payload)
		if err != nil {
			panic(err)
		}
		counter++
		//if counter > 100 {
		//	quit <- struct{}{}
		//	<-done
		//	return
		//}
		//time.Sleep(1 * time.Second)
	}

}
