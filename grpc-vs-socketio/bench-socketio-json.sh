#!/usr/bin/env zsh

alias time='/usr/bin/time -f "\nCPU: %U sec\tReal: %e sec\tRAM: %M KB"'
function bytes() {
  last=$(cat /proc/net/dev | grep lo | xargs | cut -d ' ' -f 2)
}

killall server
cd json-socketio
cd client && go build
cd ../server && go build && time ./server &
cd ../client
for n in 10000.40 10000.4000 100.4000000; do
  count=$(echo $n | cut -d '.' -f 1)
  payload=$(echo $n | cut -d '.' -f 2)
  bytes
  before=$last
  echo BENCHMARK $n
  time ./client -a 127.0.0.1:15001 -c $count -p $payload
  bytes
  after=$last
  echo $(ruby -e "p ($after-$before)/$count") bytes/rpc 
  echo 
done 
killall server
