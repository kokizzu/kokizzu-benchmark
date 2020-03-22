#!/usr/bin/env zsh

alias time='/usr/bin/time -f "\nCPU: %Us\tReal: %es\tRAM: %MKB"'
cd protobuf
cd client && go build
cd ../server && go build && time ./server &
cd ../client
echo '128'
time ./client -a 127.0.0.1:15001 -c 10000 -p 128
echo '1024'
time ./client -a 127.0.0.1:15001 -c 10000 -p 1024
echo '8192'
time ./client -a 127.0.0.1:15001 -c 10000 -p 8192
killall server

