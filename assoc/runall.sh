
alias time='/usr/bin/time -f "\nCPU: %Us\tReal: %es\tRAM: %MKB"'

set -x

pypy --version
time pypy dictionary.py
time pypy dictionary.py
time pypy dictionary.py

v -version
time v -prod run map.v
time v -prod run map.v
time v -prod run map.v

nim -v
time nim r -d:release table.nim
time nim r -d:release table.nim
time nim r -d:release table.nim

nim -v
time nim r table.nim
time nim r table.nim
time nim r table.nim

nim -v 
time nim r -d:release --gc:arc table.nim
time nim r -d:release --gc:arc table.nim
time nim r -d:release --gc:arc table.nim

nim -v
time nim r -d:release --gc:orc table.nim
time nim r -d:release --gc:orc table.nim
time nim r -d:release --gc:orc table.nim

crystal --version
time crystal run hash.cr
time crystal run hash.cr
time crystal run hash.cr

dart --version
time dart map.dart
time dart map.dart
time dart map.dart

tcc --version
time tcc -run uthash.c
time tcc -run uthash.c
time tcc -run uthash.c

go version
time go run map.go
time go run map.go
time go run map.go

java --version
time java hash_map.java
time java hash_map.java
time java hash_map.java

julia --version
time julia dict.jl
time julia dict.jl
time julia dict.jl
