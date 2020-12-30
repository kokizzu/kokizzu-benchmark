##############################
# scomb.nim

import strutils

proc newGap(gap : int) : int =
    result = (gap * 10) div 13
    if result == 9 or result == 10:
        return 11
    if result < 1:
        return 1

proc combSort(a : var seq[string], len : int) = 
    var gap = len
    var swapped = true
    while gap>1 or swapped:
        swapped = false
        gap = newGap(gap)
        for i in 0..<(len-gap):
            if a[i] > a[i+gap]:
                swapped = true
                swap(a[i],a[i+gap])


const N = 10000000

var arr = newSeq[string](N)

for z in 0..<N:
    arr[z]=intToStr(N-z)

combSort(arr,N)

for z in 1..<N:
    if arr[z]<arr[z-1]:
        stdout.write "!"
