// comb.cpp
#include<cstdlib>
#include<cstdio>
#include<algorithm>
using namespace std;
int newGap(int gap){
    gap /= 1.3;
    if(gap == 9 || gap == 10) gap = 11;
    if(gap < 1) return 1;
    return gap;
}
void combSort(int a[], int len){
    int gap = len;
    bool swapped;
    do {
        swapped = false;
        gap = newGap(gap);
        for(int i=0; i < len-gap; ++i) {
            if(a[i] > a[i+gap]) {
                swapped = true;
                swap(a[i], a[i+gap]);
            }
        }
    } while(gap > 1 || swapped);
}
const int N = 10000000;
int main() {
    int* arr = new int[N];
    for(int z=0;z<N;++z) arr[z] = rand();    
    combSort(arr,N);
    for(int z=1;z<N;++z) if(arr[z]<arr[z-1]) printf("!");
    delete[] arr;
}


