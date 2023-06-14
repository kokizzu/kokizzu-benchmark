package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	"github.com/kpango/fastime"
	"golang.org/x/sync/errgroup"
)

func main() {
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run go-gorilla-client.go localhost:8888")
		os.Exit(1)
	}

	const clientCount = 10_000

	u := url.URL{Scheme: "ws", Host: os.Args[1], Path: "/ws"}

	eg := &errgroup.Group{}
	reqCounter := int64(0)
	totalLatencyNs := int64(0)
	maxLatencyNs := int64(0)
	start := time.Now()
	for z := 0; z < clientCount; z++ {
		eg.Go(func() error {
			i := 0

			c, res, err := websocket.DefaultDialer.Dial(u.String(), nil)
			if err != nil {
				bReason, _ := io.ReadAll(res.Body)
				log.Fatalf("dial: %v, reason: %v\n", err, string(bReason))
			}
			defer c.Close()

			for {
				{
					i++
					begin := fastime.Now()
					request := fmt.Sprintf("hello %v", i)
					err := c.WriteMessage(websocket.BinaryMessage, []byte(request))
					if err != nil {
						log.Fatalf("write: %v", err)
						return err
					}

					receiveType, response, err := c.ReadMessage()
					if err != nil {
						log.Println("ReadMessage failed:", err)
						return err
					}
					if receiveType != websocket.BinaryMessage {
						log.Println("received type != websocket.BinaryMessage")
						return err

					}

					atomic.AddInt64(&reqCounter, 1)
					latency := fastime.Since(begin).Nanoseconds()
					atomic.AddInt64(&totalLatencyNs, latency)
					for {
						maxLat := atomic.LoadInt64(&maxLatencyNs)
						if maxLat >= latency {
							break
						}
						atomic.CompareAndSwapInt64(&maxLatencyNs, maxLat, latency)
					}

					if string(response) != request {
						log.Printf("'%v' != '%v'", len(response), len(request))
						return err
					}

					//log.Println("success echo websocket.BinaryMessage:", request)
				}

				{
					begin := fastime.Now()

					i++
					request := fmt.Sprintf("hello %v", i)
					err := c.WriteMessage(websocket.TextMessage, []byte(request))
					if err != nil {
						log.Fatalf("write: %v", err)
						return err
					}

					receiveType, response, err := c.ReadMessage()
					if err != nil {
						log.Println("ReadMessage failed:", err)
						return err
					}
					if receiveType != websocket.TextMessage {
						log.Printf("received type(%d) != websocket.TextMessage(%d)\n", receiveType, websocket.TextMessage)
						return err

					}

					atomic.AddInt64(&reqCounter, 1)
					latency := fastime.Since(begin).Nanoseconds()
					atomic.AddInt64(&totalLatencyNs, latency)
					for {
						maxLat := atomic.LoadInt64(&maxLatencyNs)
						if maxLat >= latency {
							break
						}
						atomic.CompareAndSwapInt64(&maxLatencyNs, maxLat, latency)
					}

					if string(response) != request {
						log.Printf("'%v' != '%v'", len(response), len(request))
						return nil
					}

					//log.Println("success echo websocket.TextMessage  :", request)
				}
				time.Sleep(100 * time.Millisecond)
			}
			return nil
		})
	}

	eg.Go(func() error {
		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <-ticker.C:
				elapsedSec := time.Since(start).Seconds()
				totalReq := float64(atomic.LoadInt64(&reqCounter))
				avgLatency := float64(atomic.LoadInt64(&totalLatencyNs)) / totalReq / 1_000_000
				maxLatency := float64(atomic.LoadInt64(&maxLatencyNs)) / 1_000_000
				rps := totalReq / elapsedSec
				fmt.Printf("\rrps: %.2f avg/max latency = %.2fms/%.2fms elapsed %.1fs\n", rps, avgLatency, maxLatency, elapsedSec)
			}
		}
		return nil
	})

	eg.Wait()
}
