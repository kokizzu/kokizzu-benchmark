
# Simple Echo WebSocket Benchmark

This is a simple echo websocket benchmark, bun/nodejs (uwebsocket) vs go+nbio. 

benchmark plan:
- 10k connection
- echo something every 10ms
- print the rps

## setup

```
# go
go mod tidy
go run go-nbio-server.go

# bun
curl -fsSL https://bun.sh/install | bash # for macOS, Linux, and WSL
bun bun-ws.js

# node
bun install uNetworking/uWebSockets.js#v20.27.0
node node-uws.js

# benchmark
go run go-gorilla-client.go localhost:8888
```

## result

```
go 1.20.5 nbio 1.3.16

rps: 253713.94 avg/max latency = 12.56ms/257.27ms elapsed 1.2
rps: 365052.46 avg/max latency = 13.24ms/257.27ms elapsed 2.2
rps: 403104.54 avg/max latency = 13.67ms/257.27ms elapsed 3.1
rps: 423652.09 avg/max latency = 13.70ms/257.27ms elapsed 4.2
rps: 437158.96 avg/max latency = 13.73ms/257.27ms elapsed 5.1
rps: 447771.98 avg/max latency = 13.68ms/257.27ms elapsed 6.1
rps: 456172.38 avg/max latency = 13.66ms/257.27ms elapsed 7.2
rps: 461172.48 avg/max latency = 13.61ms/257.27ms elapsed 8.2
rps: 465496.54 avg/max latency = 13.58ms/257.27ms elapsed 9.1
rps: 469874.77 avg/max latency = 13.55ms/257.27ms elapsed 10.2

134 (116-163) MB RAM
14-15 core usage

---

bun 0.6.9

rps: 91396.93 avg/max latency = 5.09ms/366.15ms elapsed 1.3
rps: 104604.91 avg/max latency = 12.34ms/366.15ms elapsed 2.2
rps: 114923.59 avg/max latency = 16.67ms/366.15ms elapsed 3.2
rps: 117396.76 avg/max latency = 20.56ms/366.15ms elapsed 4.2
rps: 118688.09 avg/max latency = 25.10ms/366.15ms elapsed 5.2
rps: 120622.57 avg/max latency = 28.70ms/366.15ms elapsed 6.2
rps: 121481.34 avg/max latency = 31.56ms/366.15ms elapsed 7.2
rps: 121616.55 avg/max latency = 34.50ms/366.15ms elapsed 8.2
rps: 122385.94 avg/max latency = 37.89ms/366.15ms elapsed 9.2
rps: 123589.91 avg/max latency = 40.67ms/366.15ms elapsed 10.2

55 MB RAM
1.02 core usage

---

node 18.16.0

rps: 56705.35 avg/max latency = 46.29ms/295.33ms elapsed 1.3
rps: 78918.92 avg/max latency = 58.86ms/295.33ms elapsed 2.3
rps: 93022.63 avg/max latency = 65.17ms/295.33ms elapsed 3.3
rps: 99804.09 avg/max latency = 68.52ms/295.33ms elapsed 4.3
rps: 104115.35 avg/max latency = 70.17ms/295.33ms elapsed 5.3
rps: 107737.82 avg/max latency = 70.92ms/295.33ms elapsed 6.3
rps: 109919.20 avg/max latency = 71.64ms/295.33ms elapsed 7.3
rps: 111882.40 avg/max latency = 72.00ms/295.33ms elapsed 8.3
rps: 113203.59 avg/max latency = 72.46ms/295.33ms elapsed 9.3
rps: 114449.91 avg/max latency = 72.62ms/295.33ms elapsed 10.3

59 MB RAM
1.0 core usage
```



## optional

to increase connection limit, see: https://github.com/lesismal/go-websocket-benchmark
