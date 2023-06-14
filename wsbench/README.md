
# Simple Echo WebSocket Benchmar

This is a simple echo websocket benchmark, bun (uwebsocket) vs go+nbio. 

benchmark plan:
- 10k connection
- echo something every 10ms
- print the rps

## setup

```
go mod tidy
go run go-nbio-server.go
go run go-nbio-client.go

curl -fsSL https://bun.sh/install | bash # for macOS, Linux, and WSL
bun bun-ws.js
#bun install uNetworking/uWebSockets.js#v20.27.0
```

## result

```
go+nbio

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

bun+uwebsocket
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
```



## optional

to increase connection limit, see: https://github.com/lesismal/go-websocket-benchmark
