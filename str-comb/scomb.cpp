//////////////////////////////
// scomb.cpp
#include<cstdio>
#include<string>
using namespace std;
int newGap(int gap){
    gap /= 1.3;
    if(gap == 9 || gap == 10) return 11;
    if(gap < 1) return 1;
    return gap;
}
void combSort(string a[], int len){
    int gap = len;
    bool swapped;
    do {
        swapped = false;
        gap = newGap(gap);
        for(int i=0; i < len-gap; ++i) {
            if(a[i] > a[i+gap]) {
                swapped = true;
                a[i].swap(a[i+gap]);
            }
        }
    } while(gap > 1 || swapped);
}
const int N = 10000000;
int main() {
    string* arr = new string[N];
    for(int z=0;z<N;++z) arr[z] = to_string(N-z);    
    combSort(arr,N);
    for(int z=1;z<N;++z) if(arr[z]<arr[z-1]) printf("!");
    delete[] arr;
}
