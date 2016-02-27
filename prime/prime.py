res = [3]
last = 3
while(True):
	last += 2
	prime = True
	for v in res:
		if v*v > last: break
		if last%v == 0:
			prime = False
			break
	if prime:
		res.append(last)
		if len(res)%100000 == 0: print(last)
		if last > 9999999: break
