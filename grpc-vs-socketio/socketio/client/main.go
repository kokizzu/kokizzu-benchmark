package main

import (
	"flag"
	"fmt"
	"sync"
	"time"

	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
)

var (
	host    = flag.String("a", "127.0.0.1", "the grpc server host")
	port    = flag.Int("o", 15001, "the grpc server port")
	count   = flag.Int("c", 10000, "the count of records")
	payload = flag.Int("p", 128, "the playload of the requests")
)

type Message struct {
	Text string `json:"text"`
}

func do(wg *sync.WaitGroup) {
	defer wg.Done()

	c, err := gosocketio.Dial(gosocketio.GetUrl(*host, *port, false), transport.GetDefaultWebsocketTransport())
	if err != nil {
		panic(err)
	}
	defer c.Close()

	// prepare the payload
	data := make([]byte, *payload)
	for i := 0; i < *payload; i++ {
		data[i] = 'a'
	}
	s := string(data)

	start := time.Now()

	i := 0
	for {
		i++

		if _, err := c.Ack("/message", Message{s}, time.Second*2); err != nil {
			panic(err)
		}
		if *count == i {
			break
		}
	}

	elapsed := float64(time.Since(start).Nanoseconds()) / 1000000.0
	fmt.Printf("rpc average duration: %.2f ms\n", elapsed/float64(i))
}

func main() {
	flag.Parse()

	var wg sync.WaitGroup
	wg.Add(1)
	go do(&wg)

	wg.Wait()
}
