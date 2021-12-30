
alias time='/usr/bin/time -f "\nCPU: %Us\tReal: %es\tRAM: %MKB"'

set -x

pypy --version
time pypy comb.py
time pypy comb.py
time pypy comb.py

v -version
time v -prod run comb.v
#time v -prod run comb.v
#time v -prod run comb.v

nim -v
time nim r -d:release comb.nim
#time nim r -d:release comb.nim
#time nim r -d:release comb.nim

nim -v
time nim r comb.nim
#time nim r comb.nim
#time nim r comb.nim

nim -v 
time nim r -d:release --gc:arc comb.nim
#time nim r -d:release --gc:arc comb.nim
#time nim r -d:release --gc:arc comb.nim

#nim -v
#time nim r -d:release --gc:orc comb.nim
#time nim r -d:release --gc:orc comb.nim
#time nim r -d:release --gc:orc comb.nim

crystal --version
time crystal run comb.cr
#time crystal run comb.cr
#time crystal run comb.cr

crystal --version
time crystal run --release comb.cr
#time crystal run --release comb.cr
#time crystal run --release comb.cr

dart --version
time dart comb.dart
time dart comb.dart
time dart comb.dart

tcc --version
time tcc -run comb.c
time tcc -run comb.c
time tcc -run comb.c

go version
time go run comb.go
time go run comb.go
time go run comb.go

java --version
time java comb.java
time java comb.java
time java comb.java

node --version
time node comb.js
time node comb.js
time node comb.js

pypy --version
time pypy comb.py
time pypy comb.py
time pypy comb.py

php --version
time php comb.php
time php comb.php
time php comb.php

ruby --version
time ruby comb.rb
#time ruby comb.rb
#time ruby comb.rb

python3 --version
time python3 comb.py
#time python3 comb.py
#time python3 comb.py

luajit -version
time luajit comb.lua
time luajit comb.lua
time luajit comb.lua
