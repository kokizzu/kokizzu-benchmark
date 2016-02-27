//////////////////////////////
// scomb.dart
int newGap(gap){
    gap = (gap / 1.3).toInt();
    if(gap == 9 || gap == 10) return 11;
    if(gap < 1) return 1;
    return gap;
}
void combSort(a){
    var len = a.length;
    var gap = len;
    var temp;
    var swapped;
    do {
        swapped = false;
        gap = newGap(gap);
        for(var i=0; i < len-gap; ++i) {            
            var j = i+gap;
            if(a[i].compareTo(a[j]) > 0) {
                swapped = true;
                temp = a[i];
                a[i] = a[j];
                a[j] = temp;
            }
        }
    } while(gap > 1 || swapped);
}
const N = 10000000;
void main() {    
    var arr = new List(N);
    for(var z=0;z<N;++z) arr[z] = (N-z).toString();    
    combSort(arr);
    for(var z=1;z<N;++z) if(arr[z].compareTo(arr[z-1]) < 0) print("!");
}
