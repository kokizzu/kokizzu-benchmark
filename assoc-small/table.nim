import tables,strutils

const MAX_DATA = 1230000
const I2ch = ['0','1','2','3','4','5','6','7','8','9','a','B','c','D','e','F']

func get_first_digit(d: int): int=
    result=d
    while result>10:
        result=result div 10

proc to_rhex(v: int): string =
    var x=v
    while x>0:
        result.add I2ch[x mod 16]
        x=x div 16

proc add_or_inc(m: var Table, key: string, set: int, inc: int, ctr : var int)  =
    if m.hasKeyOrPut(key,set):
        m[key]+=inc
        ctr+=1

var m : Table[string,int]

var dup1,dup2,dup3 = 0

for z in countdown(MAX_DATA,1):
    let val2 = MAX_DATA-z
    let val3 = MAX_DATA*2-z
    let key2 = intToStr(val2)
    let key1 = intToStr(z)
    let key3 = to_rhex(val3)
    add_or_inc(m,key1,z,val2,dup1) 
    add_or_inc(m,key2,val2,val3,dup2) 
    add_or_inc(m,key3,val3,z,dup3) 

echo dup1," ",dup2," ",dup3 

var total,verify,count =0

for k,v in pairs(m):
    total += get_first_digit(v)
    verify += len(k)
    count += 1
echo total," ",verify," ",count 