-- table.lua
MAX_DATA = 120000
i2ch = {'1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'B', 'c', 'D', 'e', 'F'}
i2ch[0] = '0'
get_first_digit = function(d)
	while d > 10 do
		d = d / 10
	end
	return math.floor(d)
end
to_rhex = function(v) 
	local hex = ''
	while v > 0 do
		hex = hex .. i2ch[v%16]
		v = math.floor(v / 16)
	end
	return hex
end
set_or_inc = function(m, key, set, inc) 
	if not m[key] then
		m[key] = set
		return false
	end
	m[key] = m[key] + inc		
	return true
end
m = {}
dup1, dup2, dup3 = 0, 0, 0
for z = MAX_DATA, 1, -1 do
	val2 = MAX_DATA - z
	val3 = MAX_DATA*2 - z
	key1 = tostring(z)
	key2 = tostring(val2)
	key3 = to_rhex(val3)
	if set_or_inc(m, key1, z, val2) then dup1 = dup1 + 1 end
	if set_or_inc(m, key2, val2, val3) then dup2 = dup2 + 1 end
	if set_or_inc(m, key3, val3, z) then dup3 = dup3 + 1 end
end
print(dup1, dup2, dup3)
total, verify, count = 0, 0, 0
for k, v in pairs(m) do
	total = total + get_first_digit(v)
	verify = verify + string.len(k)
	count = count + 1
end
print(total, verify, count)
