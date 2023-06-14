
# Simple Echo WebSocket Benchmark

This is a simple echo websocket benchmark, bun (uwebsocket) vs go+nbio. 

benchmark plan:
- 10k connection
- echo something every 100ms/10ms
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

sleep 100ms

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

sleep 10ms

rps: 147746.60 1.1
rps: 195998.29 2.1
rps: 221432.84 3.1
rps: 233467.68 4.1
rps: 239651.33 5.1
rps: 245445.87 6.1
rps: 249666.64 7.1
rps: 252590.65 8.1
rps: 255521.37 9.1
rps: 257248.18 10.1

134 MB RAM
14-15 core usage

---

bun 0.6.9

sleep 100ms

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

sleep 10ms

rps: 31479.49 1.1
rps: 47891.22 2.1
rps: 53222.24 3.1
rps: 55420.14 4.1
rps: 57728.77 5.1
rps: 59209.17 6.1
rps: 60032.47 7.1
rps: 60890.74 8.1
rps: 61486.55 9.1
rps: 62047.29 10.1

60 MB RAM
1.02 core usage

---

node 18.16.0

sleep 100ms

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

sleep 10ms

rps: 32614.96 1.2
rps: 42902.93 2.2
rps: 48698.52 3.2
rps: 51757.40 4.2
rps: 53558.67 5.2
rps: 54738.47 6.2
rps: 55784.62 7.2
rps: 56635.83 8.2
rps: 57446.48 9.2
rps: 57733.65 10.2

57 MB RAM
1.0 core usage
```



## optional

to increase connection limit, see: https://github.com/lesismal/go-websocket-benchmark
