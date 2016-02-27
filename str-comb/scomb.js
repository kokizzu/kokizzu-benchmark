//////////////////////////////
// scomb.js
function newGap(gap){
    gap /= 1.3;
    gap = gap|0;
    if(gap == 9 || gap == 10) return 11;
    if(gap < 1) return 1;
    return gap;
}
function combSort(a){
    var len = a.length;
    var gap = len;
    var temp;
    var swapped;
    do {
        swapped = false;
        gap = newGap(gap);
        for(var i=0; i < len-gap; ++i) {
            if(a[i] > a[i+gap]) {
                swapped = true;
                temp = a[i];
                a[i] = a[i+gap];
                a[i+gap] = temp;
            }
        }
    } while(gap > 1 || swapped);
}
var N = 10000000;
var arr = new Array(N);
for(var z=0;z<N;++z) arr[z] = '' + (N-z);    
combSort(arr);
for(var z=1;z<N;++z) if(arr[z]<arr[z-1]) console.log("!");
