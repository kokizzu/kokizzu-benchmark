# hash.rb
MAX_DATA = 1230000
I2ch     = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'B', 'c', 'D', 'e', 'F']

def get_first_digit(d)
  while d > 10
    d = d // 10
  end
  return d
end

def to_rhex(v : Int32)
  hex = String.build do |hex|
    while (v > 0)
      hex << I2ch[v % 16]
      v = v // 16
    end
  end
  return hex
end

def add_or_inc(m, key, set, inc)
  if m.has_key? key
    m[key] += inc
    return true
  end
  m[key] = set
  return false
end

m = {} of String => Int32
dup1, dup2, dup3 = 0, 0, 0
MAX_DATA.downto(1) do |z|
  val2 = MAX_DATA - z
  val3 = MAX_DATA*2 - z
  key1 = z.to_s
  key2 = val2.to_s
  key3 = to_rhex(val3)
  dup1 += 1 if add_or_inc(m, key1, z, val2)
  dup2 += 1 if add_or_inc(m, key2, val2, val3)
  dup3 += 1 if add_or_inc(m, key3, val3, z)
end
puts "#{dup1} #{dup2} #{dup3}"
total, verify, count = 0, 0, 0
m.each do |k, v|
  total += get_first_digit(v)
  verify += k.size
  count += 1
end
puts "#{total} #{verify} #{count}"
