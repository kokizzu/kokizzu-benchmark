
### TODO

1. change json-grpc folder so it implements gRPC client and server (so bench-grpc-json.sh works)
2. create a new client-server implementation using protoactor based on https://github.com/AsynkronIT/protoactor-go#networking--remoting
3. modify all other 3 implementation so it also returns sequential number (atomic.AddUint64) instead of just string
