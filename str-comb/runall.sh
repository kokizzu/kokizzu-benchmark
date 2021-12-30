
alias time='/usr/bin/time -f "\nCPU: %Us\tReal: %es\tRAM: %MKB"'

set -x

pypy --version
time pypy scomb.py
time pypy scomb.py
time pypy scomb.py

v -version
time v -prod run scomb.v
#time v -prod run scomb.v
#time v -prod run scomb.v

nim -v
time nim r -d:release scomb.nim
#time nim r -d:release scomb.nim
#time nim r -d:release scomb.nim

nim -v
time nim r scomb.nim
#time nim r scomb.nim
#time nim r scomb.nim

nim -v 
time nim r -d:release --gc:arc scomb.nim
#time nim r -d:release --gc:arc scomb.nim
#time nim r -d:release --gc:arc scomb.nim

#nim -v
#time nim r -d:release --gc:orc scomb.nim
#time nim r -d:release --gc:orc scomb.nim
#time nim r -d:release --gc:orc scomb.nim

crystal --version
time crystal run scomb.cr
#time crystal run scomb.cr
#time crystal run scomb.cr

crystal --version
time crystal run --release scomb.cr
#time crystal run --release scomb.cr
#time crystal run --release scomb.cr

dart --version
time dart scomb.dart
time dart scomb.dart
time dart scomb.dart

tcc --version
time tcc -run scomb.c
time tcc -run scomb.c
time tcc -run scomb.c

go version
time go run scomb.go
time go run scomb.go
time go run scomb.go

java --version
time java scomb.java
time java scomb.java
time java scomb.java

node --version
time node scomb.js
time node scomb.js
time node scomb.js

php --version
time php scomb.php
#time php scomb.php
#time php scomb.php

ruby --version
time ruby scomb.rb
#time ruby scomb.rb
#time ruby scomb.rb

python3 --version
time python3 scomb.py
#time python3 scomb.py
#time python3 scomb.py

luajit -version
time luajit scomb.lua
time luajit scomb.lua
time luajit scomb.lua
