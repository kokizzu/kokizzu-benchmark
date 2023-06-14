
# Simple Echo WebSocket Benchmark

This is a simple echo websocket benchmark, bun (uwebsocket) vs go+nbio. 

benchmark plan:
- 10k connection
- echo something every 10ms
- print the rps

## setup

```
go mod tidy
go run go-nbio-server.go

go run go-nbio-client.go # benchmark code using gorilla

curl -fsSL https://bun.sh/install | bash # for macOS, Linux, and WSL
bun bun-ws.js

bun install uNetworking/uWebSockets.js#v20.27.0
node node-uws.js
```

## result

```
go 1.20.5 nbio 1.3.16

rps: 43755.32 1.2
rps: 65137.37 2.2
rps: 75656.95 3.2
rps: 81256.23 4.2
rps: 84406.06 5.2
rps: 86849.70 6.2
rps: 88439.03 7.2
rps: 89629.13 8.2
rps: 90582.95 9.2
rps: 91350.43 10.2

97-107 MB RAM
4.5-4.9 core usage

---

bun 0.6.9

rps: 27355.30 1.4
rps: 39690.17 2.4
rps: 46120.58 3.4
rps: 48669.44 4.4
rps: 51504.18 5.4
rps: 52886.58 6.4
rps: 53996.81 7.4
rps: 54108.43 8.4
rps: 54448.71 9.4
rps: 54866.71 10.4

55-58 MB RAM
1.0 core usage

---

node 18.16.0

rps: 27842.13 1.1
rps: 38155.45 2.1
rps: 43114.16 3.1
rps: 45373.05 4.1
rps: 46915.13 5.1
rps: 47815.66 6.1
rps: 48494.43 7.1
rps: 49013.41 8.1
rps: 49448.08 9.1
rps: 49777.81 10.1

55 MB RAM
1.0 core usage
```



## optional

to increase connection limit, see: https://github.com/lesismal/go-websocket-benchmark
