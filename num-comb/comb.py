##############################
# comb1.py
def newGap(gap):
    gap /= 1.3
    if gap == 9 or gap == 10:
        return 11
    if gap < 1:
        return 1
    return int(gap)
def combSort(a, len):
    gap = len
    while True:
        swapped = False
        gap = newGap(gap)
        for i in range(0,len-gap):
            if a[i] > a[i+gap]:
                swapped = True
                temp = a[i]
                a[i] = a[i+gap]
                a[i+gap] = temp
        if not(gap > 1 or swapped):
            break
N = 10000000;
# comb2.py: from array import array
arr = [N-z for z in range(0,N)] # comb2.py: array('d',(N-z for z in range(0,N)))
combSort(arr,N)
for z in range(1,N):
    if arr[z]<arr[z-1]:
        print("!")
