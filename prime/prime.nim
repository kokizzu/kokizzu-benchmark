var last:int = 3
var res:seq[int] = @[]
add(res,last)
while true:
  last += 2
  var prime = true
  for v in res:
    if v*v > last: break
    if last mod v == 0:
      prime = false 
      break
  if prime:
    add(res,last)
    if len(res) mod 100000 == 0: echo last
    if last > 9999999: break
