last = 3
res = {}
res[1] = last
while true do
  last = last + 2
  prime = true
  for k,v in ipairs(res) do
   	if v*v > last then break end
   	if last % v == 0 then
      prime = false 
      break
    end
  end
  if prime then
    res[#res+1] = last 
    if #res % 100000 == 0 then print(last) end
    if last > 9999999 then break end
  end
end