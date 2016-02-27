// sel.cpp
#include<cstdlib>
#include<cstdio>
#include<algorithm>
using namespace std;
void selSort(int a[], size_t length) {
    for(size_t z=0;z<length-1;++z) {
        size_t best = z;
        for(size_t y=z+1;y<length;++y) if(a[best]>a[y]) best = y;
        if(z!=best) swap(a[best],a[z]);
    }    
}
const int N = 200000;
int main() {
    int* arr = new int[N];
    for(int z=0;z<N;++z) arr[z] = rand();
    selSort(arr,N);
    for(int z=1;z<N;++z) if(arr[z]<arr[z-1]) printf("!");
    delete[] arr;
}

