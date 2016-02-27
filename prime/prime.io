last := 3
#res := List clone append(last)
res := Sequence clone setItemType("int32") setEncoding("number")
res append(last)
loop(
   last = last + 2
   prime := true
   res foreach(i, v, 
      if(v*v > last, break)
      if(last%v == 0, 
      	 prime = false
      	 break
      )
   ) 
   if(prime,
   	  res append(last)
   	  if(res size % 100000 == 0, last println)
   	  if(last>9999999, break)
   )
)
