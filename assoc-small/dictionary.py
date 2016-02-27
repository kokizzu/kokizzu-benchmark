# dictionary.py
MAX_DATA = 120000
I2ch = ['0','1','2','3','4','5','6','7','8','9','a','B','c','D','e','F']
def get_first_digit(d):
	while(d > 10):
		d /= 10 
	return int(d)
def to_rhex(v):
	hex = ''
	v = int(v)
	while(v>0):
		hex += I2ch[v%16]
		v = int(v / 16)
	return hex
def add_or_inc(m,key,set,inc):
	if key not in m:
		m[key] = set
		return False
	m[key] += inc
	return True
m = {}
dup1, dup2, dup3 = 0, 0, 0
for z in range(MAX_DATA,0,-1):
	val2 = MAX_DATA-z
	val3 = MAX_DATA*2-z
	key2 = str(val2)
	key1 = str(z)
	key3 = to_rhex(val3)
	if add_or_inc(m,key1,z,val2): dup1 += 1
	if add_or_inc(m,key2,val2,val3): dup2 += 1
	if add_or_inc(m,key3,val3,z): dup3 += 1
print( dup1,dup2,dup3 )
total, verify, count = 0, 0, 0
for key in m:
	total += get_first_digit(m[key])
	verify += len(key)
	count += 1
print( total,verify,count )
