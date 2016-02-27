------------------------------
-- comb.lua
newGap = function(gap)
    gap = math.floor(gap / 1.3)
    if gap == 9 or gap == 10 then return 11 end
    if gap < 1 then return 1 end
    return gap
end
combSort = function(a, len)
    gap = len
    repeat
        swapped = false
        gap = newGap(gap)
        for i = 0, (len-gap-1) do
            j = i+gap
            if a[i] > a[j] then
                swapped = true                
                a[i], a[j] = a[j], a[i]
            end
        end
    until gap <= 1 and not swapped
end
N = 10000000
arr = {}
for z=0, N-1 do arr[z] = N-z end
combSort(arr,N)
for z=1, N-1 do 
    if arr[z]<arr[z-1] then print("!") end 
end
