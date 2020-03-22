// Package main implements a client for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	"kokizzu-benchmark/grpc-vs-socketio/flatbuffer/schema/users"
	"sync"
	"time"

	flatbuffers "github.com/google/flatbuffers/go"
	"google.golang.org/grpc"
)

var (
	address     = flag.String("a", "127.0.0.1:15001", "the grpc server address")
	count       = flag.Int("c", 10000, "the count of records")
	connections = flag.Int("s", 1, "the count of connections")
	payload     = flag.Int("p", 128, "the playload of the requests")
)

func do(wg *sync.WaitGroup) {
	defer wg.Done()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithCodec(flatbuffers.FlatbuffersCodec{}))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := users.NewUserServiceClient(conn)

	// prepare the payload
	data := make([]byte, *payload)
	for i := 0; i < *payload; i++ {
		data[i] = 'a'
	}
	s := string(data)

	b := flatbuffers.NewBuilder(0)
	name := b.CreateString(s)
	users.AddRequestStart(b)
	users.AddRequestAddName(b, name)
	b.Finish(users.AddRequestEnd(b))

	start := time.Now()

	i := 0
	for {
		i++

		if _, err := c.Add(context.Background(), b); err != nil {
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
	for index := 0; index < *connections; index++ {
		wg.Add(1)
		go do(&wg)
	}
	wg.Wait()
}
