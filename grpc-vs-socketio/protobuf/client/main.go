// Package main implements a client for Greeter service.
package main

import (
	"context"
	"flag"
	"sync"

	pb "kokizzu-benchmark/grpc-vs-socketio/protobuf/proto/helloworld"

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
	conn, err := grpc.Dial(*address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	// prepare the payload
	data := make([]byte, *payload)
	for i := 0; i < *payload; i++ {
		data[i] = 'a'
	}
	s := string(data)

	i := 0
	for {
		i++
		request := pb.HelloRequest{Name: s}
		if _, err := c.SayHello(context.Background(), &request); err != nil {
			panic(err)
		}

		if *count == i {
			break
		}
	}
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
