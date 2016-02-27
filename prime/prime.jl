function main() 
  res = [last=3]
  while true
     last += 2
     prime = true
     for v in res
       if v*v > last 
         break
       end
       if last % v == 0 
         prime = false 
         break
       end
     end
     if prime
       push!(res,last)
       if length(res) % 100000 == 0 
         println(last)
       end 
       if last > 9999999 
         break
       end
     end
  end
end
main()
