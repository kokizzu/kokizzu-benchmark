#!/usr/bin/env zsh

alias time='/usr/bin/time -f "\nCPU: %Us\tReal: %es\tRAM: %M KB"'
function bytes() {
  return `cat /proc/net/dev | grep lo | xargs | cut -d ' ' -f 2`
}

killall server
cd json-socketio
cd client && go build
cd ../server && go build && time ./server &
cd ../client
for payload in 128 1024 32768; do
  for count in 10000; do
    bytes
    before=$?
    echo BENCHMARK $count x $payload
    time ./client -a 127.0.0.1:15001 -c $count -p $payload
    bytes
    after=$?
    echo `expr $after - $before` bytes 
    echo 
  done
done 
killall server
